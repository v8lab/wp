package market

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

var SafeWebSocketDestroyError = fmt.Errorf("connection destroy by user")

// SafeWebSocket 安全的WebSocket封装
// 保证读取和发送操作是并发安全的，支持自定义保持alive函数
type SafeWebSocket struct {
	ws               *websocket.Conn
	listener         SafeWebSocketMessageListener
	aliveHandler     SafeWebSocketAliveHandler
	aliveInterval    time.Duration
	sendQueue        chan []byte
	lastError        error
	runningTaskSend  bool
	runningTaskRead  bool
	runningTaskAlive bool
}

type SafeWebSocketMessageListener = func(b []byte)
type SafeWebSocketAliveHandler = func()

// NewSafeWebSocket 创建安全的WebSocket实例并连接
func NewSafeWebSocket(endpoint string) (*SafeWebSocket, error) {
	ws, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		return nil, err
	}
	s := &SafeWebSocket{ws: ws, sendQueue: make(chan []byte, 1000), aliveInterval: time.Second * 60}

	go func() {
		s.runningTaskSend = true
		for s.lastError == nil {
			b := <-s.sendQueue
			if err := s.ws.WriteMessage(websocket.TextMessage, b); err != nil {
				s.lastError = err
				break
			}
		}
		s.runningTaskSend = false
	}()

	go func() {
		s.runningTaskRead = true
		for s.lastError == nil {
			_, b, err := s.ws.ReadMessage()
			if err != nil {
				s.lastError = err
				break
			}
			s.listener(b)
		}
		s.runningTaskRead = false
	}()

	go func() {
		s.runningTaskAlive = true
		for s.lastError == nil {
			if s.aliveHandler != nil {
				s.aliveHandler()
			}
			time.Sleep(s.aliveInterval)
		}
		s.runningTaskAlive = false
	}()

	return s, nil
}

// Listen 监听消息
func (s *SafeWebSocket) Listen(h SafeWebSocketMessageListener) {
	s.listener = h
}

// Send 发送消息
func (s *SafeWebSocket) Send(b []byte) {
	s.sendQueue <- b
}

// KeepAlive 设置alive周期及函数
func (s *SafeWebSocket) KeepAlive(v time.Duration, h SafeWebSocketAliveHandler) {
	s.aliveInterval = v
	s.aliveHandler = h
}

// Destroy 销毁
func (s *SafeWebSocket) Destroy() (err error) {
	s.lastError = SafeWebSocketDestroyError
	for !s.runningTaskRead && !s.runningTaskSend && !s.runningTaskAlive {
		time.Sleep(time.Millisecond * 100)
	}
	if s.ws != nil {
		err = s.ws.Close()
		s.ws = nil
	}
	s.listener = nil
	s.aliveHandler = nil
	s.sendQueue = nil
	return err
}

// Loop 进入事件循环，直到连接关闭才退出
func (s *SafeWebSocket) Loop() error {
	for s.lastError == nil {
		time.Sleep(time.Millisecond * 100)
	}
	return s.lastError
}

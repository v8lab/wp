package main

import (
	"fmt"

	"github.com/leizongmin/huobiapi"
)

func main() {
	// 创建客户端实例
	market, err := huobiapi.NewMarket()
	if err != nil {
		panic(err)
	}
	// 订阅主题
	market.Subscribe("market.eosusdt.trade.detail", func(topic string, json *huobiapi.JSON) {
		// 收到数据更新时回调
		fmt.Println(topic, json)
	})
	// 请求数据
	json, err := market.Request("market.eosusdt.detail")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(json)
	}
	// 进入阻塞等待，这样不会导致进程退出
	market.Loop()
}

package huobiapi

import (
	"fmt"
	"testing"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/stretchr/testify/assert"
	"github.com/leizongmin/huobiapi/data_type"
)

func TestNewMarket(t *testing.T) {
	m, err := NewMarket()
	assert.NoError(t, err)
	go func() {
		time.Sleep(time.Second * 10)
		m.Close()
	}()
	m.Subscribe("market.eosusdt.trade.detail", func(topic string, json *simplejson.Json) {
		fmt.Println(topic, json)
		b, err := json.Encode()
		assert.NoError(t, err)
		fmt.Println(data_type.DecodeTrade(b))
	})
	m.Loop()
}

func TestNewClient(t *testing.T) {
	c, err := NewClient("", "")
	assert.NoError(t, err)
	ret, err := c.Request("GET", "/market/history/trade", ParamsData{"symbol": "eosusdt", "size": "10"})
	assert.NoError(t, err)
	fmt.Println(ret)
}

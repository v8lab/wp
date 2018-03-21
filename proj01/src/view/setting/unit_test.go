package setting

import (
	"fmt"
	"testing"
)

func Test_Unit(t *testing.T) {

	var Unit UnitStu
	Unit.HeaderData = `{"header":{"tokenshare":["clienttestid01"],"sessionid":["12345678123456781234567812345678"]}}`
	Unit.Init()
	fmt.Println(Unit.Header)
	for k, v := range Unit.Header {
		fmt.Println(k, v)
	}
	fmt.Println(Unit.Header.Get("tokenshare"))
}

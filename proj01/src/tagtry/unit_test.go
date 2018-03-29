package tagtry

import (
	"fmt"
	"reflect"
	"testing"
	Tool "tool"
)

func Test_tag(t *testing.T) {
	Log := Tool.NewTwLogStu("%v\t%v\t\n")
	defer func() {
		fmt.Println(Log.Flush())
	}()
	var Client ClientStu
	Client.Id = "id01"
	Client.Data = "data01"

	Client_ := TagFunc(&Client)
	Log.Print("Client_", Client_)

	datatag, _ := Client_["data"]
	Log.Print("datatag", datatag)

	vCli := reflect.ValueOf(&Client)
	Log.Print("vCli string", vCli.String())

	Log.Print("-------", "")

	Elems := reflect.ValueOf(&Client).Elem()
	for i := 0; i < Elems.NumField(); i++ {
		Log.Print("Elems.Type().Field(i)", Elems.Type().Field(i))
	}

}

package tagtry

import (
	"fmt"
	"testing"
)

func Test_tag(t *testing.T) {
	var Client ClientStu
	Client.Id = "id01"
	Client.Data = "data01"

	Client_ := TagFunc(&Client)
	datatag, _ := Client_["data"]
	fmt.Println(datatag)
}

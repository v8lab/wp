package unittest

import (
	"fmt"
	"testing"

	mystu "mystu"
	type01 "mystu/type01"
)

func Test_Mystu(t *testing.T) {
	var Entry01 mystu.EntryStu
	Entry01.Init("entry01")
	Entry01.Data = "afghjk"
	fmt.Println(Entry01.ToJson())

	var Entrytype01 type01.EntryType01Stu
	Entrytype01.InitByJson(Entry01.ToJson())
	fmt.Println(Entrytype01.ToJson())

}

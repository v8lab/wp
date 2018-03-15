package myflag

import (
	"testing"
)

func Test_myflag(t *testing.T) {
	Data := "pram show -client=c001 -person=p001"
	Myflag := NewMyflagStu(Data)
	Myflag.Print()
}

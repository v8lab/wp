package a17observer

import (
	"testing"
)

func Test_Func(t *testing.T) {

	var Subject SubjectStu

	Subject.Register(&OberverStu{})
	Subject.Register(&OberverStu{})
	Subject.Register(&OberverStu{})

	Subject.Change()
}

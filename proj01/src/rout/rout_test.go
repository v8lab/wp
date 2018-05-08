package rout

import (
	"fmt"
	"testing"
	"time"
)

func Test_rout(t *testing.T) {
	times := 9
	chans := make(chan int, times)

	for i := 0; i < times; i++ {
		vv := i
		go func() {

			time.Sleep(time.Second * time.Duration(i))
			chans <- vv
		}()
	}

	for i := 0; i < times; i++ {

		fmt.Println(<-chans)

	}
}

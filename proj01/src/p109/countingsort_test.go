package p109

import (
	"fmt"
	"testing"
)

func Test_CountingSort(t *testing.T) {

	aData := []int{2, 5, 3, 0, 2, 3, 0, 3}
	bData := make([]int, len(aData))
	// k
	k := 0
	for _, v := range aData {
		if v >= k {
			k = v
		}
	}
	fmt.Println(k)
	// init cdata 0
	cData := make([]int, k+1)
	fmt.Println(cData)
	//
	for _, v := range aData {
		cData[v] = cData[v] + 1
	}

	for i := 1; i < len(cData); i++ {
		cData[i] = cData[i] + cData[i-1]
	}
	fmt.Println(cData)
	for _, v := range aData {
		bData[cData[v]-1] = v
		cData[v] = cData[v] - 1
	}
	fmt.Println(aData)
	fmt.Println(bData)
	fmt.Println(cData)
}

package code_test

import (
	"fmt"
	"testing"
	"gmr/gmr-x/offer/code"
)

func TestNum01(t *testing.T) {
	var arr = [][]int{{1,2,8,9},{2,4,9,12},{4,7,10,13}}
	// result := NumOneMethodOne(2,arr)
	result := code.Num01MethodTwo(3,arr)
	fmt.Println(result)
	
}
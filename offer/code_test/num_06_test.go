package code_test

import (
	"fmt"
	"gmr/gmr-x/offer/code"
	"testing"

)

func TestNum06(t *testing.T) {
	var arr = []int{2,3,4,1}
	min := code.Num06MethodOne(arr)
	fmt.Println(min)
}

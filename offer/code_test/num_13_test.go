package code_test

import (
	"fmt"
	"gmr/gmr-x/offer/code"
	"testing"
)

func TestNum13(t *testing.T) {
	var arr = []int{1,1,3,4,5,6,7,8}
	res := code.Num13MethodThree(arr)
	fmt.Println(res)
}

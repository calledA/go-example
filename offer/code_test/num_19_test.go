package code_test

import (
	"fmt"
	"gmr/gmr-x/offer/code"
	"testing"
)

func TestNum19(t *testing.T) {
	var arr = [][]int{{1,2,3},{2,3,4},{1,2,5}}
	fmt.Println(code.Num19MethodOne(arr))
}

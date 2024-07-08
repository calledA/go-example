package code_test

import (
	"fmt"
	"gmr/gmr-x/offer/code"
	"testing"
)

func TestNum21(t *testing.T) {
	var pushed = []int{1, 2, 3, 4}
	var popped = []int{3,4, 2, 1}
	fmt.Println(code.Num21MethodOne(pushed, popped))
}

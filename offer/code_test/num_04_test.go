package code_test

import (
	"gmr/gmr-x/ds"
	"gmr/gmr-x/offer/code"
	"testing"
)

func TestNum04(t *testing.T) {
	var preOrder = []int{1, 2, 4, 8, 5, 9, 3, 6, 7, 10}
	var midOrder = []int{4, 8, 2, 9, 5, 1, 6, 3, 7, 10}
	root := code.Num04MethodOne(preOrder, midOrder)
	ds.LDRTree(root)
}

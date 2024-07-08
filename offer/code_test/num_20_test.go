package code_test

import (
	"fmt"
	"gmr/gmr-x/offer/code"
	"testing"
)

func TestNum20(t *testing.T) {
	minStack := code.Constructor()
	minStack.Push(2)
	minStack.Push(1)
	minStack.Push(3)
	minStack.Push(3)
	minStack.Push(3)
	minStack.Push(-1)
	minStack.Pop()
	minStack.Pop()
	fmt.Println(minStack.Val)
	fmt.Println(minStack.Mini)
	fmt.Println(minStack.Min())
}

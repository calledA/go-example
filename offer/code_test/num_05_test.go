package code_test

import (
	"fmt"
	"gmr/gmr-x/offer/code"
	"testing"
)

func TestNum05(t *testing.T) {
	var array = code.ArrayStack{
		Array: []string{"11", "22", "33"},
		Size:  3,
	}
	var pop = array.Pop()
	fmt.Println(pop)
	array.Push("44")
	fmt.Println(len(array.Array))
	fmt.Println(array.Array)
}

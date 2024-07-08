package code

import (
	"math"
)

/*
定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的min函数（时间复杂度应为O（1））。
*/

type MinStack struct {
	Val  []int
	Mini []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{math.MaxInt64}}
}

func (minStack *MinStack) Push(val int) {
	minStack.Val = append(minStack.Val, val)
	top := minStack.Mini[len(minStack.Mini)-1]
	minStack.Mini = append(minStack.Mini, min(top, val))
}

func (minStack *MinStack) Pop() {
	minStack.Val = minStack.Val[:len(minStack.Val)-1]
	minStack.Mini = minStack.Mini[:len(minStack.Mini)-1]
}

func (minStack *MinStack) Top() int {
	return minStack.Val[len(minStack.Val)-1]
}

func (minStack *MinStack) Min() int {
	return minStack.Mini[len(minStack.Mini)-1]
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

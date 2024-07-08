package code

import "gmr/gmr-x/ds"

/*
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否可能为该栈的弹出顺序。
假设压入栈的所有数字均不相等。例如序列1,2,3,4,5是某栈的压入顺序，序列4,5,3,2,1是该压栈序列对应的一个弹出序列，但4,3,5,1,2就不可能是该压栈序列的弹出序列
*/

/*
解题思路 
1，本题主要考察入栈出栈的理解
2，golang slice可以很容易实现栈
3，每次pushed入栈后popped 进行比较
4，如果栈非空，且poped的当前元素和栈顶元素相等，则出栈，同时右移popped指针
*/
func Num21MethodOne(pushed, popped []int) bool {
	if len(pushed) == 0 || len(popped) == 0 || len(pushed) != len(popped) {
		return false
	}
	var stack ds.ArrayStack
	var j = 0
	for i := 0; i < len(pushed); i++ {
		stack.Push(pushed[i])
		for !stack.IsEmpty() && stack.Peek() == popped[j] {
			stack.Pop()
			j++
		}
	}
	return stack.IsEmpty() && j==len(popped)
}


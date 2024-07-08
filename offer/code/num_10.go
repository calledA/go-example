package code

/*
我们可以用2*1的小矩形横着或者竖着去覆盖更大的矩形。
请问用n个2*1的小矩形无重叠地覆盖一个2*n的大矩形，总共有多少种方法？
*/

func Num10MethodOne(num int) int {
	if num <= 2 {
		return num
	}
	var first, second, third = 1, 2, 3
	for i := 3; i <= num; i++ {
		third = first + second
		first = second
		second = third
	}
	return third
}

func Num10MethodTwo(num int) int {
	if num <= 2 {
		return num
	}
	return Num10MethodTwo(num-1) + Num10MethodTwo(num-2)
}

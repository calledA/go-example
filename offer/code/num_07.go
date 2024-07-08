package code

/*
大家都知道斐波那契数列，现在要求输入一个整数n，请你输出斐波那契数列的第n项（从0开始，第0项为0，第1项是1）。
*/

func Num07MethodOne(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	var first, second, third = 0, 1, 1
	for i := 2; i <= num; i++ {
		third = first + second
		first = second
		second = third
	}
	return third
}

func Num07MethodTwo(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 || num == 2 {
		return 1
	}
	return Num07MethodTwo(num-1) + Num07MethodTwo(num-2)
}

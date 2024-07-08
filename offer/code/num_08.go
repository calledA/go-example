package code

/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）。
*/

func Num08MethodOne(num int) int {
	if num == 1 {
		return 1
	}
	var first, second = 1, 2
	for i := 3; i <= num; i++ {
		var third = first + second
		first = second
		second = third
	}
	return second
}

func Num08MethodTwo(num int) int {
	if num == 1 {
		return 1
	}
	if num == 2 {
		return 2
	}
	return Num08MethodTwo(num-1) + Num08MethodTwo(num-2)
}

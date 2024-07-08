package code

/*
给定一个double类型的浮点数base和int类型的整数exponent。求base的exponent次方。

保证base和exponent不同时为0。不得使用库函数，同时不需要考虑大数问题，也不用考虑小数点后面0的位数。
*/

func Num12MethodOne(base float64, exponent int) float64 {
	if base == 0.0 {
		return 0.0
	}
	if exponent == 0 {
		return 1.0
	}
	var flag = false
	if exponent < 0 {
		flag = true
		exponent *= -1
	}
	var result = base
	for i := 2; i <= exponent; i++ {
		result *= base
	}
	if flag {
		return 1.0 / result
	}
	return result
}

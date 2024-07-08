package code

import "fmt"

/*
输入一个整数，输出该数32位二进制表示中1的个数。其中负数用补码表示。
*/

func Num11MethodOne(num uint32) int {
	sum := 0
	bin := fmt.Sprintf("%b", num)
	for _, value := range bin {
		if value == '1' {
			sum++
		}
	}
	return sum
}

//优解，位运算
func Num11MethodTwo(num uint32) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

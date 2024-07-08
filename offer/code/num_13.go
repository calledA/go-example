package code

/*
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前半部分，所有的偶数位于数组的后半部分，并保证奇数和奇数，偶数和偶数之间的相对位置不变。
*/

func Num13MethodOne(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	var newArr = make([]int, 0)
	for _, value := range arr {
		if (value & 1) == 1 {
			newArr = append(newArr, value)
		}
	}
	for _, value := range arr {
		if (value & 1) == 0 {
			newArr = append(newArr, value)
		}
	}
	return newArr
}

//不能保证相对位置不变 leetcode版
func Num13MethodTwo(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	var slow, fast = 0, 0
	for fast < len(arr) {
		if arr[fast]&1 != 0 {
			arr[slow], arr[fast] = arr[fast], arr[slow]
			slow++
		}
		fast++
	}
	return arr
}

//O(N) 优解
func Num13MethodThree(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	var newArr = make([]int, 0)
	var oddIndex, evenIndex = 0, 0
	for _, value := range arr {
		if value&1 == 1 {
			arr[oddIndex] = value
			oddIndex++
		} else {
			newArr = append(newArr, value)
			evenIndex++
		}
	}
	for i := 0; i < evenIndex; i++ {
		arr[oddIndex+i] = newArr[i]
	}
	return arr
}

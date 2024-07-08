package code

/*
在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
*/

//从右上角逐渐逼近
func Num01MethodOne(target int, arr [][]int) bool {
	if len(arr) == 0 || len(arr[0]) == 0 {
		return false
	}
	var row, col = len(arr), len(arr[0])
	var w, h = col - 1, 0
	for w > 0 && h < row {
		if arr[h][w] > target {
			w--
		} else if arr[h][w] < target {
			h++
		} else {
			return true
		}
	}
	return false
}

//因为是递增数组，可以使用二分查找
func Num01MethodTwo(target int, arr [][]int) bool {
	if len(arr) == 0 || len(arr[0]) == 0 {
		return false
	}
	for i := 0; i < len(arr); i++ {
		if hasFound(target, arr[i]) {
			return true
		}
	}
	return false
}

func hasFound(target int, arr []int) bool {
	var start, end = 0, len(arr) - 1
	for start+1 < end {
		var mid = start + (end-start)/2
		if arr[mid] == target {
			return true
		} else if arr[mid] > target {
			end = mid
		} else {
			start = mid
		}
	}
	if arr[start] == target || arr[end] == target {
		return true
	}
	return false
}

package code

func Num29MethodOne(nums []int, k int) []int {
	var res = nums[:k]
	for i := k; i < len(nums); i++ {
		max := SelectSort(res)
		if res[max] > nums[i] {
			res[max], nums[i] = nums[i], res[max]
		}
	}
	return res
}

func SelectSort(data []int) int {
	max := 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[max] {
			max = i
		}
	}
	return max
}

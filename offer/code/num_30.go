package code

func Num30MethodOne(nums []int) int {
	var sum, maxSum int
	for i := 0; i < len(nums); i++ {
		sum = 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

package code

func Num28MethodOne(nums []int) int {
	hash := make(map[int]int,len(nums))
	temp := len(nums) / 2
	for _, val := range nums {
		if _,ok := hash[val];ok {
			hash[val] += 1
			if hash[val] > temp {
				return val
			}
		} else {
			hash[val] = 1
		}
	}
	return 0
}

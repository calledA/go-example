package code

/*
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个非递减排序的数组的一个旋转，输出旋转数组的最小元素。例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1。
NOTE：给出的所有元素都大于0，若数组大小为0，请返回0。
*/

func Num06MethodOne(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	for index, _ := range arr {
		if index == len(arr)-1 {
			break
		}
		if arr[index] > arr[index+1] {
			return arr[index+1]
		}
	}
	return arr[0]
}

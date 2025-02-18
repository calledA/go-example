package code

/*
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字，例如，如果输入如下4 X 4矩阵： 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 则依次打印出数字1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10.
*/

func Num19MethodOne(matrix [][]int) []int {
	var res []int
	row := len(matrix)
	col := len(matrix[0])

	if row == 0 {
		return res
	}

	res = make([]int, row*col)
	total, rowStart, colStart := 0, 0, 0
	rowEnd, colEnd := row-1, col-1

	horizonDirection, verticalDirection := true, true

	for total < col+row {
		if horizonDirection {
			for i := colStart; i <= colEnd; i++ {
				res[total] = matrix[rowStart][i]
				total += 1
			}
			rowStart += 1
		} else {
			for i := colEnd; i >= colStart; i-- {
				res[total] = matrix[rowEnd][i]
				total += 1
			}
			rowEnd -= 1
		}
		horizonDirection = !horizonDirection

		if rowStart > rowEnd {
			break
		}

		if verticalDirection {
			for i := rowStart; i <= rowEnd; i++ {
				res[total] = matrix[i][colEnd]
				total += 1
			}
			colEnd -= 1
		} else {
			for i := rowEnd; i >= rowStart; i-- {
				res[total] = matrix[i][colStart]
				total += 1
			}
			colStart += 1
		}
		verticalDirection = !verticalDirection
		if colStart > colEnd {
			break
		}
	}
	return res
}

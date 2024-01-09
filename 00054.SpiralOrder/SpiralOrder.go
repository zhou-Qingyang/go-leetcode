package _0054_SpiralOrder

func spiralOrder(matrix [][]int) []int {
	if len(matrix[0]) == 0 {
		return []int{}
	}
	n, m := len(matrix), len(matrix[0])

	left, right := 0, m-1
	top, bottom := 0, n-1
	res := []int{}

	for left <= right && top <= bottom {
		//从左到右  最上边
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		// 从上到下  可定是最右边
		for j := top; j <= bottom; j++ {
			res = append(res, matrix[j][right])
		}
		right--

		// 从右到左输出一行
		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		// 从下到上输出一列
		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}

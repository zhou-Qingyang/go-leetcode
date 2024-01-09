package _0059_GenerateMatrix

func generateMatrix(num int) [][]int {

	if num == 0 {
		return [][]int{}
	}

	left, right := 0, num-1
	top, bottom := 0, num-1

	matrix := make([][]int, num)
	for k, _ := range matrix {
		matrix[k] = make([]int, num)
	}
	k := 1
	for left <= right && top <= bottom {
		// 从左向右开始赋值
		for i := left; i <= right; i++ {
			matrix[top][i] = k
			k++
		}
		top++

		// 上至下
		for i := top; i <= bottom; i++ {
			matrix[i][right] = k
			k++
		}
		right--

		// 从右到左输出一行
		if top <= bottom {
			for i := right; i >= left; i-- {
				matrix[bottom][i] = k
				k++
			}
			bottom--
		}

		// 从下到上输出一列
		if left <= right {
			for i := bottom; i >= top; i-- {
				matrix[i][left] = k
				k++
			}
			left++
		}
	}
	return matrix
}

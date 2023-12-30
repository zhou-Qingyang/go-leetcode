package _0118_Generate

func generate(maxRows int) [][]int {
	if maxRows == 0 {
		return [][]int{}
	}
	//res := make([][]int, maxRows)
	//
	res := [][]int{}
	res = append(res, []int{0})
	for i := 1; i < maxRows; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; i++ {
			row[j] = res[i-1][j-1] + res[i-1][j]
		}
		res = append(res, row)
	}
	return res
}

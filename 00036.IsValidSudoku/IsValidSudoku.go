package _0036_IsValidSudoku

func isValidSudoku(target [][]byte) bool {

	for i := 0; i < 9; i++ {
		isUsed := make(map[byte]bool, 9)
		for j := 0; j < 9; j++ {
			if target[i][j] == '.' {
				continue
			}
			if isUsed[target[i][j]] {
				return false
			}
			isUsed[target[i][j]] = true
		}
	}

	// 判断每一列是否合法
	for j := 0; j < 9; j++ {
		isUsed := make(map[byte]bool, 9)
		for i := 0; i < 9; i++ {
			if target[i][j] == '.' {
				continue
			}
			if isUsed[target[i][j]] {
				return false
			}
			isUsed[target[i][j]] = true
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			isUsed := make(map[byte]bool, 9)
			for k := 0; k < 3; k++ {
				for f := 0; f < 3; f++ {
					if target[i+k][j+f] == '.' {
						continue
					}
					if isUsed[target[i+k][j+f]] {
						return false
					}
					isUsed[target[i+k][j+f]] = true
				}
			}
		}
	}
	return true
}

package Game_354_of_week

func sumOfSquares(target []int) int {
	ans := 0
	n := len(target)
	for key, val := range target {
		index := key + 1
		if n%index == 0 {
			ans += val * val
		}
	}
	return ans
}

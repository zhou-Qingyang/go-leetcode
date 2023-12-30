package _0042_Trap

func trap(height []int) int {
	pre_max := make([]int, len(height))
	pre_max[0] = height[0]
	for i := 1; i < len(height); i++ {
		pre_max[i] = max(pre_max[i-1], height[i])
	}

	suf_max := make([]int, len(height))
	suf_max[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		suf_max[i] = max(suf_max[i+1], height[i])
	}

	sum := 0
	for i, v := range height {
		ans := min(pre_max[i], suf_max[i]) - v
		sum += ans
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

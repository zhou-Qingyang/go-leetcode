package leetcode

func jump(nums []int) int {
	n := len(nums)
	maxPos := 0
	end := 0
	steps := 0

	for i := 0; i < n-1; i++ {
		maxPos = max(maxPos, i+nums[i])
		if i == end {
			steps++
			end = maxPos
			if end >= n-1 {
				break
			}
		}
	}

	return steps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

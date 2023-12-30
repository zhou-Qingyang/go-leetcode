package _0041_FirstMissingPositive

func firstMissingPositive(nums []int) int {

	n := len(nums) // 获取数组的长度

	// 第一次遍历数组，将每个正整数尽可能地放在正确的位置上
	for i := 0; i < n; i++ {
		// 当前元素 nums[i] 大于 0、小于等于数组长度，并且不在正确的位置上时
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			// 交换当前元素和应该在的位置上的元素
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1

}

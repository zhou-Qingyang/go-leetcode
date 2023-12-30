package leetcode

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	// 测试用例 [1,2,2,2,3,4,6]
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}

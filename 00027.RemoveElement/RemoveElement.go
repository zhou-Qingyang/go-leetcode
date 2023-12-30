package _0027_RemoveElement

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	// 测试用例 [1,2,2,2,3,3,4,6]
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow + 1
}

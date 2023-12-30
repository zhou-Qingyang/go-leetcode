package _0035_SearchInsert

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		} else if target < nums[i] {
			return i - 1
		}
	}
	return len(nums)
}

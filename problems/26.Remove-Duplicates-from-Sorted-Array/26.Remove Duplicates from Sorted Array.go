package leetcode

// 1,2,3,4
func removeDuplicates(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	n := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[n] {
			n++
			nums[n] = nums[i]
		}
	}
	return n + 1
}

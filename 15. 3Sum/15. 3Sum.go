func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			sum := 0 - nums[i]
			left := i + 1
			right := len(nums) - 1
			for left < right {
				if nums[left]+nums[right] == sum {
					res = append(res, []int{nums[i], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--

					}
					left++
					right--
				} else if nums[left]+nums[right] > sum {
					right--
				} else {
					left++
				}
			}
		}
	}

	return res
}
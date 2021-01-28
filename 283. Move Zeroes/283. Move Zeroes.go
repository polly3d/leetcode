func moveZeroes(nums []int) {
	zeroCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
		} else {
			nums[i-zeroCount], nums[i] = nums[i], nums[i-zeroCount]
		}
	}

}
func sortedSquares(nums []int) []int {
	n := len(nums)
	i := n - 1
	j := 0
	res := make([]int, n)
	for k := n - 1; k >= 0; k-- {
		if abs(nums[i]) > abs(nums[j]) {
			res[k] = nums[i] * nums[i]
			i--
		} else {
			res[k] = nums[j] * nums[j]
			j++
		}
	}
	return res
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
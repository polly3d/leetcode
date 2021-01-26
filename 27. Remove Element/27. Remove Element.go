func removeElement(nums []int, val int) int {
	clen := 0
	for _, v := range nums {
		if v != val {
			nums[clen] = v
			clen++
		}
	}
	return clen
}
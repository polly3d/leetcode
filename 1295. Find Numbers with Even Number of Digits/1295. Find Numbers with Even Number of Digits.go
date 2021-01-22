func findNumbers(nums []int) int {
	count := 0
	for _, n := range nums {
		if digitsNum(n)%2 == 0 {
			count++
		}
	}
	return count
}

func digitsNum(n int) int {
	count := 0
	for n != 0 {
		n = n / 10
		count++
	}
	return count
}
func findDisappearedNumbers(nums []int) []int {
	res := make([]int, len(nums))
	for _, v := range nums {
		res[v-1] = v
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if res[i] == 0 {
			res[j] = i + 1
			j++
		}
	}
	return res[:j]
}

func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	for _, v := range nums {
		if nums[abs(v)-1] > 0 {
			nums[abs(v)-1] *= -1
		}
	}
	for i, v := range nums {
		if v > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}
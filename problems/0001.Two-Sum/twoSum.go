package leetcode

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for ai, num := range nums {
		b := target - num
		if bi, ok := m[b]; ok {
			return []int{ai, bi}
		} else {
			m[num] = ai
		}
	}
	return nil
}

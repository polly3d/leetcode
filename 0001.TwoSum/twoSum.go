package leetcode

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, n := range nums {
		b, ok := m[target-n]
		if ok {
			return []int{b, i}
		}
		m[n] = i
	}
	return nil
}

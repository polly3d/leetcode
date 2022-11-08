package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	res := 0
	min := math.MaxInt
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			distance := abs(sum - target)
			if distance < min {
				res = sum
				min = distance
			}
			if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return res
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

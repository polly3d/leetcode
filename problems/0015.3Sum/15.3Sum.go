package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	target := 0
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			if nums[i]+nums[j]+nums[k] < target {
				j++
				for j < len(nums) && nums[j-1] == nums[j] {
					j++
				}
			} else if nums[i]+nums[j]+nums[k] > target {
				k--
				for k > i && nums[k+1] == nums[k] {
					k--
				}
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j-1] == nums[j] {
					j++
				}
				for j < k && nums[k+1] == nums[k] {
					k--
				}
			}
		}
	}
	return res
}

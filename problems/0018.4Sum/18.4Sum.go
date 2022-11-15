package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}
			start, end := j+1, len(nums)-1
			for start < end {
				if nums[i]+nums[j]+nums[start]+nums[end] < target {
					start++
					for start < len(nums) && nums[start-1] == nums[start] {
						start++
					}
				} else if nums[i]+nums[j]+nums[start]+nums[end] > target {
					end--
					for end > j && nums[end+1] == nums[end] {
						end--
					}
				} else {
					res = append(res, []int{nums[i], nums[j], nums[start], nums[end]})
					start++
					end--
					for start < end && nums[start-1] == nums[start] {
						start++
					}
					for start < end && nums[end+1] == nums[end] {
						end--
					}
				}
			}
		}
	}
	return res
}

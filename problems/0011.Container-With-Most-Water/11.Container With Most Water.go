package leetcode

import "leetcode/utils"

func maxArea(height []int) int {
	start, end, max := 0, len(height)-1, 0
	for start < len(height) && end >= 0 {
		w := end - start
		h := utils.Min(height[end], height[start])
		max = utils.Max(w*h, max)
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	return max
}

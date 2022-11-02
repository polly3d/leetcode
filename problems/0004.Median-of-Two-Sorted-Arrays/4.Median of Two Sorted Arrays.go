package leetcode

import (
	"leetcode/utils"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	if len1 > len2 {
		return findMedianSortedArrays(nums2, nums1)
	}
	if len1 == 0 {
		return float64((nums2[(len2-1)/2] + nums2[len2/2]) / 2)
	}
	total := len1 + len2

	start, end := 0, len1
	cut1, cut2 := 0, 0
	for start <= end {
		cut1 = (start + end) / 2
		cut2 = (total+1)/2 - cut1
		l1 := math.MinInt
		if cut1 > 0 {
			l1 = nums1[cut1-1]
		}
		r1 := math.MaxInt
		if cut1 < len1 {
			r1 = nums1[cut1]
		}
		l2 := math.MinInt
		if cut2 > 0 {
			l2 = nums2[cut2-1]
		}
		r2 := math.MaxInt
		if cut2 < len2 {
			r2 = nums2[cut2]
		}
		if l1 > r2 {
			end = cut1 - 1
		} else if l2 > r1 {
			start = cut1 + 1
		} else {
			if total%2 == 0 {
				return float64(utils.Max(l1, l2)+utils.Min(r1, r2)) / 2.0
			} else {
				return float64(utils.Max(l1, l2))
			}
		}
	}

	return -1
}

package leetcode

import (
	"leetcode/utils"
)

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	length := 0
	for i := 0; i < len(s); i++ {
		m := map[byte]bool{}
		m[s[i]] = true
		length++
		maxLength = utils.Max(maxLength, length)
		for j := i + 1; j < len(s); j++ {
			r := m[s[j]]
			if r {
				length = 0
				break
			} else {
				m[s[j]] = true
				length++
				maxLength = utils.Max(maxLength, length)
			}
		}
		length = 0
	}
	return maxLength
}

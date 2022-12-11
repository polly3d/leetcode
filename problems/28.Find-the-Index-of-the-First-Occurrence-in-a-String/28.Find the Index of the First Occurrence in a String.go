package leetcode

func strStr(haystack string, needle string) int {
	i := -1
	for {
		i++
		j := -1
		for {
			j++
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}

			if haystack[i+j] != needle[j] {
				break
			}
		}
	}
}

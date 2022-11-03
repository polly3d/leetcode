package leetcode

// Dp
func longestPalindrome(s string) string {
	res, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j] && (j-i < 3 || dp[i+1][j-1]))
			if dp[i][j] && (res == "" || j+1-i > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

// Slide window
func longestPalindrome2(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(s, i, i, res)
		res = maxPalindrome(s, i, i+1, res)
	}
	return res
}

func maxPalindrome(s string, i, j int, longest string) string {
	res := ""
	// Palindrome defination
	for i >= 0 && j < len(s) && s[i] == s[j] {
		res = s[i : j+1]
		i--
		j++
	}
	if len(res) > len(longest) {
		return res
	}
	return longest
}

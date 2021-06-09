func longestPalindrome(s string) string {
	ls := ""
	for i := 0; i <= len(s); i++ {
		for j := i; j <= len(s); j++ {
			ts := string(s[i:j])
			if isPalidromic(ts) {
				if len(ts) > len(ls) {
					ls = ts
				}
			}
		}
	}
	return ls
}

func isPalidromic(s string) bool {
	sl := len(s)
	for i := 0; i < sl/2; i++ {
		if s[i] != s[sl-i-1] {
			return false
		}
	}
	return true
}
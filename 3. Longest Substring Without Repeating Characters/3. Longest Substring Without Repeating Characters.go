func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	i, max := 0, 0
	for i < len(s) {
		count := 0
		m := map[string]bool{}
		for j := i; j < len(s); j++ {
			r := string(s[j])
			if _, ex := m[r]; ex {
				break
			} else {
				m[r] = true
				count++
				if count > max {
					max = count
				}
			}
		}
		i++
	}
	return max
}
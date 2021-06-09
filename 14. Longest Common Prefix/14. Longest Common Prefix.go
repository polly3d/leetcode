func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	first := strs[0]
	if len(strs) < 2 {
		return first
	}
	preFix := ""
	for r, s := range first {
		c := true
		for i := 1; i < len(strs); i++ {
			str := strs[i]
			if r > len(str)-1 {
				c = false
				continue
			}
			if string(s) != string(str[r]) {
				c = false
				break
			}
		}
		if c {
			preFix += string(s)
		} else {
			break
		}
	}
	return preFix
}
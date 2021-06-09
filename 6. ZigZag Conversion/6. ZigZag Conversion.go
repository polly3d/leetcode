func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	ss := make([][]string, numRows)
	i := 0
	count := 0
	down := true
	for i < len(s) {
		fmt.Println(count)
		ss[count] = append(ss[count], string(s[i]))
		if down {
			count++
			if count >= numRows-1 {
				down = false
			}
		} else {
			count--
			if count <= 0 {
				down = true
			}
		}
		i++
	}
	fmt.Println(ss)
	res := ""
	for _, rs := range ss {
		res += strings.Join(rs, "")
	}
	return res
}
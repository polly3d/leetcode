package leetcode

func convert(s string, numRows int) string {
	down, up, ss := 0, numRows-2, make([][]byte, numRows)
	for i := 0; i < len(s); {
		if down < numRows {
			ss[down] = append(ss[down], s[i])
			down++
			i++
		} else if up > 0 {
			ss[up] = append(ss[up], s[i])
			up--
			i++
		} else {
			down = 0
			up = numRows - 2
		}
	}
	res := []byte{}
	for _, r := range ss {
		res = append(res, r...)
	}
	return string(res)
}

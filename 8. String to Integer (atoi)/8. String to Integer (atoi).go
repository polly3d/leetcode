func myAtoi(s string) int {
	res := 0
	if s == "" {
		return 0
	}
	//Remove leading whitespace
	i := 0
	for i < len(s)-1 {
		if string(s[i]) != " " {
			break
		}
		i++
	}
	s = s[i:]
	//Negative or positive
	sign := 1
	if string(s[0]) == "-" || string(s[0]) == "+" {
		if string(s[0]) == "-" {
			sign = -1
		}
		s = s[1:]
	}

	//Read digit
	for _, r := range s {
		if !(r >= '0' && r <= '9') {
			break
		}
		res = res*10 + int(r) - '0'
		if sign*res > math.MaxInt32 {
			return math.MaxInt32
		} else if sign*res < math.MinInt32 {
			return math.MinInt32
		}

	}

	return res * sign
}
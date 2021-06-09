func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	res := 0
	i := 0
	for i < len(s) {
		next := i + 1
		r := 0
		if next < len(s) {
			ts := string(s[i]) + string(s[next])
			r = romanMap(ts)
		}
		if r != 0 {
			i++
			res += r
		} else {
			res += romanMap(string(s[i]))
		}
		i++
	}
	return res
}

func romanMap(s string) int {
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	case "IV":
		return 4
	case "IX":
		return 9
	case "XL":
		return 40
	case "XC":
		return 90
	case "CD":
		return 400
	case "CM":
		return 900
	}
	return 0
}
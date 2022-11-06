package leetcode

func romanToInt(s string) int {
	romans := map[string]int{"M": 1000, "CM": 900, "D": 500, "CD": 400, "C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1}
	res := 0
	for i := 0; i < len(s); {
		j := i + 1
		if j < len(s) {
			n, exist := romans[s[i:j+1]]
			if exist {
				res += n
				i += 2
			} else {
				n := romans[s[i:i+1]]
				res += n
				i++
			}
		} else {
			n := romans[s[i:i+1]]
			res += n
			i++
		}
	}
	return res
}

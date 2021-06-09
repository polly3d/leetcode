func intToRoman(num int) string {
	v := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	m := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	s := ""
	i := 0
	for num != 0 {
		for v[i] > num {
			i++
		}
		num -= v[i]
		s += m[i]
	}
	return s
}

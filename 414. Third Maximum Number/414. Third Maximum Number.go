func thirdMax(nums []int) int {
	m, s, t := math.MinInt64, math.MinInt64, math.MinInt64
	for _, v := range nums {
		if v == m || v == s || v == t {
			continue
		}
		switch {
		case v > m:
			m, s, t = v, m, s
		case v > s:
			s, t = v, s
		case v > t:
			t = v
		}
	}
	if t == math.MinInt64 {
		return m
	}
	return t
}
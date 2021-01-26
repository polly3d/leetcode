func checkIfExist(arr []int) bool {
	m := map[int]int{}
	for i, v := range arr {
		d := v * 2
		_, ok := m[d]
		if ok {
			return true
		}
		d = v / 2
		_, ok = m[d]
		if ok && v%2 == 0 {
			return true
		}
		m[v] = i
	}
	return false
}
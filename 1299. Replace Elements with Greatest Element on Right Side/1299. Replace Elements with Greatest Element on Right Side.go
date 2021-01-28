func replaceElements(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		arr[i] = findGreatest(arr[i+1:])
	}
	return arr
}

func findGreatest(arr []int) int {
	g := -1
	for _, v := range arr {
		if g < v {
			g = v
		}
	}
	return g
}
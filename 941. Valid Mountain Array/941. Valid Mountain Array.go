func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	n := len(arr)

	i := 0
	for i < n-1 {
		if arr[i] >= arr[i+1] {
			break
		}
		i++
	}
	if i == 0 {
		return false
	}
	j := i + 1
	for j < n {
		if arr[j] >= arr[j-1] {
			return false
		}
		j++
	}
	if j == i+1 {
		return false
	}
	return true
}
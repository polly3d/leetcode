func sortArrayByParity(A []int) []int {
	i, j := 0, len(A)-1
	for i != j {
		if A[i]%2 != 0 {
			if A[j]%2 == 0 {
				A[i], A[j] = A[j], A[i]
			} else {
				j--
			}
		} else {
			i++
		}
	}
	return A
}
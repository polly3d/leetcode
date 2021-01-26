func duplicateZeros(arr []int) {
	t := []int{}
	for _, v := range arr {
		t = append(t, v)
		if v == 0 {
			t = append(t, v)
		}
	}
	for i, _ := range arr {
		arr[i] = t[i]
	}
}

func duplicateZeros(arr []int) {
	zeroCount := 0
	for _, v := range arr {
		if v == 0 {
			zeroCount++
		}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			if i+zeroCount < len(arr) {
				arr[i+zeroCount] = 0
			}
			zeroCount--
		}
		if i+zeroCount < len(arr) {
			arr[i+zeroCount] = arr[i]
		}
	}
}
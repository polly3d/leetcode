func findMedianSortedArrays(a []int, b []int) float64 {
	l1 := len(a)
	l2 := len(b)

	if l1 > l2 {
		return findMedianSortedArrays(b, a)
	}

	if l1 == 0 {
		return float64(b[l2/2]+b[(l2-1)/2]) / 2.0
	}

	ml := (l1 + l2 + 1) / 2
	start, end := 0, l1
	for start <= end {
		ca := (start + end + 1) / 2
		cb := ml - ca
		la := math.MinInt64
		if ca != 0 {
			la = a[ca-1]
		}
		lb := math.MinInt64
		if cb != 0 {
			lb = b[cb-1]
		}
		ra := math.MaxInt64
		if ca != l1 {
			ra = a[ca]
		}
		rb := math.MaxInt64
		if cb != l2 {
			rb = b[cb]
		}
		if la > rb {
			end = ca - 1
		} else if lb > ra {
			start = ca + 1
		} else {
			if (l1+l2)%2 == 0 {
				return float64(max(la, lb)+min(ra, rb)) / 2.0
			}
			return float64(max(la, lb))
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
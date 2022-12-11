package utils

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

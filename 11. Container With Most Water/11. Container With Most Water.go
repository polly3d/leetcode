func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			w := j - i
			h := min(height[i], height[j])
			area := w * h
			if area > max {
				max = area
			}
		}
	}
	return max
}

func maxArea(height []int) int {
	max := 0
	l := 0
	r := len(height) - 1
	for l < r {
		w := r - l
		h := min(height[l], height[r])
		area := w * h
		if area > max {
			max = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
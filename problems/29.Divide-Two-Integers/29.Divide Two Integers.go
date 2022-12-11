package leetcode

import (
	"leetcode/utils"
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == math.MaxInt32 && divisor == -1 {
		return math.MaxInt32
	}

	result := 0
	times := 0
	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}

	dividend = utils.Abs(dividend)
	divisor = utils.Abs(divisor)

	for dividend >= divisor {
		cur := dividend - divisor<<times
		if cur >= 0 {
			result += 1 << times
			times++
			dividend = cur
		} else {
			times = 0
		}
	}

	return result * sign
}

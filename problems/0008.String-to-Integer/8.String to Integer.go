package leetcode

import (
	"math"
)

func myAtoi(s string) int {
	rs, sign := "", 0
	for _, v := range s {
		if v == ' ' && len(rs) == 0 {
			continue
		}
		if v == '-' && len(rs) == 0 {
			sign = -1
			continue
		}
		if v == '+' && len(rs) == 0 {
			sign = 0
			continue
		}
		if '0' <= v && v <= '9' {
			rs += string(v)
		}
	}
	l := len(rs) - 1
	num := 0
	numMap := map[rune]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
	for _, n := range rs {
		num += numMap[n] * int(math.Pow10(l))
		l--
	}
	if sign == -1 {
		num *= -1
	}
	if num < math.MinInt32 {
		return math.MinInt32
	}
	if num > math.MaxInt32 {
		return math.MaxInt32
	}
	return num
}

package leetcode

func intToRoman(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := ""
	for num != 0 {
		for i := 0; i < len(nums); i++ {
			if num >= nums[i] {
				res += romans[i]
				num = num - nums[i]
				break
			}
		}
	}
	return res
}

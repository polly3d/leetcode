package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	reverse := 0
	for tmp != 0 {
		reverse = reverse*10 + tmp%10
		tmp = tmp / 10
	}
	return reverse == x
}

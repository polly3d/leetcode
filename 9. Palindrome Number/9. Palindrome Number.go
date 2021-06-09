func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	t := x
	r := 0
	for t != 0 {
		r = r*10 + t%10
		t /= 10
	}
	return r == x
}
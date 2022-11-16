package leetcode

func isValid(s string) bool {
	stack := []string{}
	pairs := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}

	for _, v := range s {
		_, ok := pairs[string(v)]
		if ok {
			stack = append(stack, string(v))
		} else {
			if len(stack) > 0 {
				// pop and match
				top := stack[len(stack)-1]
				m := pairs[top]
				if m != string(v) {
					return false
				}
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		}
	}

	return true
}

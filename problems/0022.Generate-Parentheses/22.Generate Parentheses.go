package leetcode

func generateParenthesis(n int) []string {
	res := []string{}
	dfs(n, 0, "", &res)
	return res
}

func dfs(left, right int, s string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, s)
		return
	}
	if left > 0 {
		dfs(left-1, right+1, s+"(", res)
	}
	if right > 0 {
		dfs(left, right-1, s+")", res)
	}
}

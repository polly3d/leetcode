package leetcode

var letters = []string{
	"",
	"",
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

func letterCombinations(digits string) []string {
	res := []string{}
	if digits == "" {
		return res
	}
	dfs(digits, 0, "", &res)
	return res
}

func dfs(digitals string, index int, s string, res *[]string) {
	if index == len(digitals) {
		*res = append(*res, s)
		return
	}
	num := digitals[index]
	letter := letters[num-'0']
	for i := 0; i < len(letter); i++ {
		dfs(digitals, index+1, s+string(letter[i]), res)
	}
}

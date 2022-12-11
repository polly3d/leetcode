package leetcode

func findSubstring(s string, words []string) []int {
	res := []int{}
	if len(s) <= 0 || len(words) <= 0 {
		return res
	}
	wordCount := len(words)
	wordLength := len(words[0])
	wordsMap := map[string]int{}
	for _, w := range words {
		wordsMap[w]++
	}
	for i := 0; i < len(s)-wordCount*wordLength; i++ {
		see := map[string]int{}
		j := 0
		for j < wordCount {
			word := s[i+j*wordLength : i+(j+1)*wordLength]
			if _, ok := wordsMap[word]; ok {
				see[word]++
				if see[word] > wordsMap[word] {
					break
				}
			} else {
				break
			}
			j++
		}
		if j == wordCount {
			res = append(res, i)
		}
	}
	return res
}

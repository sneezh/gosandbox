package main

func partition(s string) [][]string {
	result := [][]string{}
	DFS(&result, []string{}, 0, s)
	return result
}

func DFS(result *[][]string, cur []string, offset int, s string) {
	if len(s) == offset {
		*result = append(*result, append([]string{}, cur...))
		return
	}

	for i := offset; i < len(s); i++ {
		subStr := s[offset : i+1]
		if isPalindrome(subStr) {
			cur = append(cur, subStr)
			DFS(result, cur, i+1, s)
			cur = cur[:len(cur)-1]
		}
	}
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

package LeetCode

var MAP = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

// letterCombinations
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var ans, s = make([]string, 0), make([]byte, 0)
	dfs(0, &s, digits, &ans)

	return ans
}

func dfs(step int, s *[]byte, digits string, ans *[]string) {
	if step == len(digits) {
		*ans = append(*ans, string(*s))
		return
	}

	var num = digits[step] - '0'
	for _, c := range MAP[num] {
		*s = append(*s, byte(c))
		dfs(step+1, s, digits, ans)
		*s = (*s)[:len(*s)-1]
	}
}

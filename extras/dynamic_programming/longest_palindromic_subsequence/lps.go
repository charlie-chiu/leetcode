package longest_palindromic_subsequence

// basic dp:
// recursive solution with memoization
var table [][]int

func longestPalindromeSubseq(s string) int {
	// table with initial value 0
	table = make([][]int, len(s))
	for i := range table {
		table[i] = make([]int, len(s))
	}

	return lpsHelper(s, 0, len(s)-1)
}

func lpsHelper(s string, start, end int) int {
	if table[start][end] > 0 {
		return table[start][end]
	}

	//every single character are palindrome
	if start == end {
		table[start][end] = 1
		return 1
	}

	if end-start == 1 && s[start] == s[end] {
		table[start][end] = 2
		return 2
	}

	if s[start] == s[end] {
		table[start][end] = lpsHelper(s, start+1, end-1) + 2
		return table[start][end]
	}

	table[start][end] = max(lpsHelper(s, start+1, end), lpsHelper(s, start, end-1))
	return table[start][end]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

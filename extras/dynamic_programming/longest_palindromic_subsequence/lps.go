package longest_palindromic_subsequence

func longestPalindromeSubseq(s string) int {
	// initial table
	// half of table is useless
	// space complexity: O(n^2) n = length of s
	seqLen := len(s)
	var table = make([][]int, seqLen)
	for i := range table {
		table[i] = make([]int, seqLen)
	}

	// fill the table
	for subseqLength := 1; subseqLength <= seqLen; subseqLength++ {
		for start := 0; start+subseqLength <= seqLen; start++ {
			end := start + subseqLength - 1
			if start == end {
				table[start][end] = 1
			} else if subseqLength == 2 && s[start] == s[end] {
				table[start][end] = 2
			} else if s[start] == s[end] {
				table[start][end] = table[start+1][end-1] + 2
			} else {
				table[start][end] = max(table[start+1][end], table[start][end-1])
			}
		}
	}

	return table[0][seqLen-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// basic dp:
// recursive solution with memoization

//var table [][]int
//func longestPalindromeSubseq(s string) int {
//	// table with initial value 0
//	table = make([][]int, len(s))
//	for i := range table {
//		table[i] = make([]int, len(s))
//	}
//
//	return lpsHelper(s, 0, len(s)-1)
//}
//
//func lpsHelper(s string, start, end int) int {
//	if table[start][end] > 0 {
//		return table[start][end]
//	}
//
//	//every single character are palindrome
//	if start == end {
//		table[start][end] = 1
//		return 1
//	}
//
//	if end-start == 1 && s[start] == s[end] {
//		table[start][end] = 2
//		return 2
//	}
//
//	if s[start] == s[end] {
//		table[start][end] = lpsHelper(s, start+1, end-1) + 2
//		return table[start][end]
//	}
//
//	table[start][end] = max(lpsHelper(s, start+1, end), lpsHelper(s, start, end-1))
//	return table[start][end]
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}

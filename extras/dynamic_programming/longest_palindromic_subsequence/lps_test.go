package longest_palindromic_subsequence

import "testing"

type TestCase struct {
	sequence string
	answer   int
}

func TestLPS(t *testing.T) {
	var TestCases = []TestCase{
		{"A", 1},
		{"BB", 2},
		{"cbbd", 2},
		{"BBBAB", 4},
		{"GEEKSFORGEEKS", 5},

		// TLE
		{"euazbipzncptldueeuechubrcourfpftcebikrxhybkymimgvldiwqvkszfycvqyvtiwfckexmowcxztkfyzqovbtmzpxojfofbvwnncajvrvdbvjhcrameamcfmcoxryjukhpljwszknhiypvyskmsujkuggpztltpgoczafmfelahqwjbhxtjmebnymdyxoeodqmvkxittxjnlltmoobsgzdfhismogqfpfhvqnxeuosjqqalvwhsidgiavcatjjgeztrjuoixxxoznklcxolgpuktirmduxdywwlbikaqkqajzbsjvdgjcnbtfksqhquiwnwflkldgdrqrnwmshdpykicozfowmumzeuznolmgjlltypyufpzjpuvucmesnnrwppheizkapovoloneaxpfinaontwtdqsdvzmqlgkdxlbeguackbdkftzbnynmcejtwudocemcfnuzbttcoew", 159},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := longestPalindromeSubseq(tc.sequence)
			if got != tc.answer {
				t.Logf("wrong answer")
				t.Logf("input: %q", tc.sequence)
				t.Logf("want: %d, got %d", tc.answer, got)
				t.Fail()
			}
		})
	}
}

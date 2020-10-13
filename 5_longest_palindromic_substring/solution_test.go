package longest_palindromic_substring

import "testing"

type TestCase struct {
	input, answer string
}

func TestLongestPalindrome(t *testing.T) {
	var TestCases = []TestCase{
		{"a", "a"},
		{"ac", "a"}, // or c ?
		{"ccc", "ccc"},
		{"cbbd", "bb"},
		{"babad", "bab"},
		{"aacabdkacaa", "aca"},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := longestPalindrome(tc.input)
			if got != tc.answer {
				t.Errorf("input: %q, want %q, got %q", tc.input, tc.answer, got)
			}
		})
	}
}

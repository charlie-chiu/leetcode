package isomorphic_strings

import "testing"

type TestCase struct {
	s, t   string
	answer bool
}

func TestIsomorphic(t *testing.T) {
	var TestCases = []TestCase{
		// note : You may assume both s and t have the same length.
		{"ab", "aa", false},
		{"aba", "baa", false},
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := isIsomorphic(tc.s, tc.t)
			if got != tc.answer {
				t.Logf("wrong answer")
				t.Logf("input: %q , %q", tc.s, tc.t)
				t.Logf("  got: %v", got)
				t.Fail()
			}
		})
	}
}

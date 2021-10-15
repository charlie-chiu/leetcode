package valid_anagram

import "testing"

type TestCase struct {
	s, t   string
	answer bool
}

func TestIsAnagram(t *testing.T) {
	var TestCases = []TestCase{
		{"anagram", "nagaram", true},
		{"a", "b", false},
		{"b", "bbb", false},
		{"rat", "car", false},
		{"aacc", "ccac", false},

		//Follow up: What if the inputs contain Unicode characters?
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := isAnagram(tc.t, tc.s)
			if got != tc.answer {
				t.Logf("wrong answer at input %q & %q", tc.s, tc.t)
				t.Logf("want: %v, got: %v", tc.answer, got)
				t.Fail()
			}
		})
	}
}

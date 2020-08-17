package implement_strstr

import "testing"

type TestCase struct {
	haystack, needle string
	answer           int
}

func TestStrStr(t *testing.T) {
	testCases := []TestCase{
		{"world", "", 0},
		{"gopher", "r", 5},
		{"aloha", "ll", -1},
		{"hello", "ll", 2},
		{"ll", "hello", -1},
	}

	for _, tc := range testCases {
		got := strStr(tc.haystack, tc.needle)

		if got != tc.answer {
			t.Errorf("wrong answer:\n haystack:%q, needle:%q\n  want:%v\n  got:%v", tc.haystack, tc.needle, tc.answer, got)
		}
	}
}

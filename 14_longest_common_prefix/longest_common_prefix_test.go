package longest_common_prefix

import "testing"

type TestCase struct {
	input  []string
	answer string
}

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []TestCase{
		{[]string{}, ""},
		{[]string{"abc"}, "abc"},
		{[]string{"abc", "db", "third"}, ""},
		{[]string{"abc", "db"}, ""},
		{[]string{"flower", "flow", "flight", "flat"}, "fl"},
	}

	for _, tc := range testCases {
		got := longestCommonPrefix(tc.input)

		if got != tc.answer {
			t.Errorf("wrong answer:\n input:%v\n  want:%v\n  got:%v", tc.input, tc.answer, got)
		}
	}
}

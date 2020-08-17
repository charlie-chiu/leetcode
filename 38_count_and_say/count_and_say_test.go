package count_and_say

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input  int
	answer string
}

func TestCountAndSay(t *testing.T) {
	testCases := []TestCase{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
		{6, "312211"},
		{7, "13112221"},
		{8, "1113213211"},
		{9, "31131211131221"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input %d", tc.input), func(t *testing.T) {
			got := countAndSay(tc.input)

			if got != tc.answer {
				t.Errorf("wrong answer:\n want:%v\n  got:%v", tc.answer, got)
			}
		})
	}
}

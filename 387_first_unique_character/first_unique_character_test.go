package first_unique_character

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	solution int
}

func TestProblem(t *testing.T) {
	var TestCases = []TestCase{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aabbcc", -1},
	}

	for _, testCase := range TestCases {
		t.Run(fmt.Sprintf("input: %q", testCase.input), func(t *testing.T) {
			got := firstUniqChar(testCase.input)
			if got != testCase.solution {
				t.Errorf("wrong answer, got: %d", testCase.solution)
			}
		})
	}
}

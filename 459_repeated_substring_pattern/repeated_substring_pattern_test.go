package repeated_substring_pattern

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input  string
	output bool
}

func TestSolution(t *testing.T) {
	var TestCases = []TestCase{
		{"aa", true},
		{"aaa", true},
		{"abab", true},
		{"aba", false},
		{"abcabcabc", true},
		{"abcabcabcabc", true},
	}

	for _, testCase := range TestCases {
		testName := fmt.Sprintf("test input %q got %v", testCase.input, testCase.output)
		t.Run(testName, func(t *testing.T) {
			if repeatedSubstringPattern(testCase.input) != testCase.output {
				t.Errorf("got wrong answer on input %q", testCase.input)
			}
		})
	}
}

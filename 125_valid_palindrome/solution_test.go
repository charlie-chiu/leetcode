package palindrome

import (
	"testing"
)

type TestCase struct {
	input string
	want  bool
}

func TestSolution(t *testing.T) {
	testCases := []TestCase{
		{"marmot", false},
		{"ab@a", true},
		{"a.", true},
		{"", true},
		{"tacocat", true},
		{"taco cat", true},
		{"A man, a plan, a canal: Panama", true},
	}

	for _, tc := range testCases {

		got := solution(tc.input)

		if got != tc.want {
			t.Errorf("result not equal on %q, got %v", tc.input, got)
		}
	}
}

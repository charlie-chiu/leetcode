package palindrome

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input  string
	answer bool
}

func TestIsPalindrome(t *testing.T) {
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
		testName := fmt.Sprintf("input:%q,answer:%v", tc.input, tc.answer)
		t.Run(testName, func(t *testing.T) {
			got := isPalindrome(tc.input)

			if got != tc.answer {
				t.Errorf("wrong answer, input %q, got %v", tc.input, got)
			}
		})
	}
}

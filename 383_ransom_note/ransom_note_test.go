package ransom_note

import (
	"fmt"
	"testing"
)

type TestCase struct {
	ransomNote, magazine string
	answer               bool
}

func TestCanConstruct(t *testing.T) {
	var TestCases = []TestCase{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
		{"aba", "aab", true},
	}

	for _, testCase := range TestCases {
		testName := fmt.Sprintf("note=%q,magazine=%q,want %v", testCase.ransomNote, testCase.magazine, testCase.answer)
		t.Run(testName, func(t *testing.T) {
			got := canConstruct(testCase.ransomNote, testCase.magazine)
			if got != testCase.answer {
				t.Errorf("got %v", got)
			}
		})
	}
}

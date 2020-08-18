package reverse_vowels

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input, output string
}

func TestReverseVowels(t *testing.T) {
	var TestCases = []TestCase{
		{"hello", "holle"},
		{"leetcode", "leotcede"},
		{"aA", "Aa"},
	}

	for _, tc := range TestCases {
		t.Run(fmt.Sprintf("reverse %q to %q", tc.input, tc.output), func(t *testing.T) {
			got := reverseVowels(tc.input)

			if got != tc.output {
				t.Errorf("got %q", got)
			}
		})
	}
}

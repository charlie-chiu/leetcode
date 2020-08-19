package integer_to_roman

import (
	"fmt"
	"testing"
)

type TestCase struct {
	arabic int
	roman  string
}

func TestIntToRoman(t *testing.T) {
	var TestCases = []TestCase{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{9, "IX"},
		{10, "X"},
		{20, "XX"},
		{58, "LVIII"},
		{95, "XCV"},
		{1994, "MCMXCIV"},
	}

	for _, testCase := range TestCases {
		testName := fmt.Sprintf("%d gets converted to %q", testCase.arabic, testCase.roman)
		t.Run(testName, func(t *testing.T) {
			got := intToRoman(testCase.arabic)
			if got != testCase.roman {
				t.Errorf("got %q, want %q", got, testCase.roman)
			}
		})
	}
}

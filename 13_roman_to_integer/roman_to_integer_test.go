package roman_to_integer

import (
	"fmt"
	"testing"
)

type TestCase struct {
	arabic int
	roman  string
}

func TestRomanToInteger(t *testing.T) {
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
		{58, "LVIII"},
		{95, "XCV"},
		{1994, "MCMXCIV"},
	}

	for _, testCase := range TestCases {
		testName := fmt.Sprintf("%q gets converted to %d", testCase.roman, testCase.arabic)
		t.Run(testName, func(t *testing.T) {
			got := romanToInt(testCase.roman)

			if got != testCase.arabic {
				t.Errorf("got %d, want %d", got, testCase.arabic)
			}
		})
	}
}

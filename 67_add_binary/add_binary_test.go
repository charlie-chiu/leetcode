package add_binary

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a, b, sum string
}

func TestAddBinary(t *testing.T) {
	var TestCases = []TestCase{
		{"10", "1", "11"},
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
		{"1", "1011", "1100"},
		{
			// overflows uint64
			"10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
			"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011",
			"110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000",
		},
	}

	for _, tc := range TestCases {
		testName := fmt.Sprintf("%q+%q=%q", tc.a, tc.b, tc.sum)
		t.Run(testName, func(t *testing.T) {
			got := addBinary(tc.a, tc.b)
			if got != tc.sum {
				t.Errorf("wrong answer, got %q", got)
			}
		})
	}
}

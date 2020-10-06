package zigzag_conversion

import "testing"

type TestCase struct {
	input   string
	numRows int
	answer  string
}

func TestZigZag(t *testing.T) {
	var TestCases = []TestCase{
		{"A", 1, "A"},
		{"A", 2, "A"},
		{"AB", 1, "AB"},
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := convert(tc.input, tc.numRows)
			if got != tc.answer {
				t.Errorf("wrong answer")
				t.Logf("want %s got %s", tc.answer, got)
			}
		})
	}
}

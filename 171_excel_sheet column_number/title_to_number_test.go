package excel_sheet_column_number

import "testing"

type TestCase struct {
	Column string
	Number int
}

func TestSolution(t *testing.T) {
	var TestCases = []TestCase{
		{"A", 1},
		{"B", 2},
		{"C", 3},
		{"AB", 28},
		{"ZY", 701},
	}

	for _, testCase := range TestCases {
		got := solution(testCase.Column)
		if got != testCase.Number {
			t.Errorf("assert Column %s covert to %d, got %d", testCase.Column, testCase.Number, got)
		}
	}
}

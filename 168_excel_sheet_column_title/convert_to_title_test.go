package excel_sheet_column_title

import "testing"

type TestCase struct {
	title  string
	number int
}

func TestSolution(t *testing.T) {
	var TestCases = []TestCase{
		{"A", 1},
		{"B", 2},
		{"C", 3},
		{"AB", 28},
		{"ZY", 701},
		{"YYZ", 17576},
		{"ABBT", 19000},
	}

	for _, testCase := range TestCases {
		got := solution(testCase.number)
		if got != testCase.title {
			t.Errorf("assert number %d covert to %s, got %s", testCase.number, testCase.title, got)
		}
	}
}

package missing_number

import "testing"

type TestCase struct {
	input  []int
	output int
}

func TestSolution(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{0}, 1},
		{[]int{3, 0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}

	for _, testCase := range TestCases {
		got := solution(testCase.input)

		if got != testCase.output {
			t.Errorf("assert failed on input %d, want %d, got %d", testCase.input, testCase.output, got)
		}
	}
}

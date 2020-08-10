package contains_duplicate

import (
	"testing"
)

type TestCase struct {
	input  []int
	output bool
}

func TestSolution(t *testing.T) {

	testCases := []TestCase{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, tc := range testCases {
		got := solution(tc.input)
		if got != tc.output {
			t.Errorf("wrong answer: input: %v, output: %v, want: %v", tc.input, got, tc.output)
		}
	}
}

package factorial_trailling_zeroes

import (
	"testing"
)

type TestCase struct {
	input, output int
}

func TestSolution(t *testing.T) {
	var TestCases = []TestCase{
		{700, 174},
		//{3, 0},
		//{5, 1},
		//{10, 2},
		//{14, 2},
		//{15, 3},
		//{19, 3},
		//{20, 4},
		//{21, 4},
		//{22, 4},
		//{23, 4},
		//{24, 4},
		//{25, 6},
		//{29, 6},
		//{30, 7},
		//{40, 9},
		//{50, 12},
		//{60, 14},
		//{70, 16},
		//{80, 19},
		//{90, 21},
		//{100, 24}, //?
		//{200, 49}, //?
		//{300, 74}, //?
		//{30000000, 7499998}, //?
	}

	for _, testCase := range TestCases {
		got := solution(testCase.input)

		if got != testCase.output {
			t.Errorf("assert failed on input %d, want %d, got %d", testCase.input, testCase.output, got)
		}
	}
}

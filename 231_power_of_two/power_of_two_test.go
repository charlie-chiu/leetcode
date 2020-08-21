package power_of_two

import (
	"testing"
)

type TestCase struct {
	input  int
	output bool
}

func TestIsPowerOfTwo(t *testing.T) {
	var TestCases = []TestCase{
		{1, true},
		{16, true},
		{20, false},
		{32767, false},
		{32768, true},
	}

	for _, testCase := range TestCases {
		got := isPowerOfTwo(testCase.input)
		if got != testCase.output {
			t.Errorf("test %d is power of two: want %v, got %v", testCase.input, testCase.output, got)
		}
	}
}

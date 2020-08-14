package third_maximum

import (
	"testing"
)

type TestCase struct {
	input  []int
	output int
}

func TestMoveZeros(t *testing.T) {
	testCases := []TestCase{
		{[]int{3, 2, 1}, 1},
		{[]int{1, 2, 3}, 1},
		{[]int{1, 2}, 2},
		{[]int{2, 2, 3, 1}, 1},
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 1, 2}, 2},
	}

	for _, tc := range testCases {
		input := make([]int, len(tc.input))
		copy(input, tc.input)

		got := thirdMax(input)

		if got != tc.output {
			t.Errorf("wrong answer:\ninput: %v\n want: %v\n  got: %v", tc.input, tc.output, got)
		}
	}
}

package move_zeros

import (
	"reflect"
	"testing"
)

type TestCase struct {
	input, output []int
}

func TestMoveZeros(t *testing.T) {
	testCases := []TestCase{
		{[]int{1, 1, 0}, []int{1, 1, 0}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
		{[]int{1, 0, 1}, []int{1, 1, 0}},
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	}

	for _, tc := range testCases {
		input := make([]int, len(tc.input))
		copy(input, tc.input)
		//input := tc.input
		moveZeroes(input)

		if !reflect.DeepEqual(input, tc.output) {
			t.Errorf("wrong answer:\ninput: %v\n want: %v\n  got: %v", tc.input, tc.output, input)
		}
	}
}

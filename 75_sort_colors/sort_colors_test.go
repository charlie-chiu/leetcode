package sort_colors

import (
	"reflect"
	"testing"
)

type TestCase struct {
	input, answer []int
}

func TestSortColors(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{0, 1, 2}, []int{0, 1, 2}},

		{[]int{1, 2, 0}, []int{0, 1, 2}},

		{[]int{2, 0, 1}, []int{0, 1, 2}},

		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},

		{[]int{1, 0, 0, 2, 2, 1}, []int{0, 0, 1, 1, 2, 2}},

		{[]int{2, 0, 2, 1, 1, 0, 0, 2, 2, 2}, []int{0, 0, 0, 1, 1, 2, 2, 2, 2, 2}},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			input := make([]int, len(tc.input))
			copy(input, tc.input)

			sortColors(tc.input)

			if !reflect.DeepEqual(tc.input, tc.answer) {
				t.Errorf("wrong answer on input %v", input)
				t.Logf("want %v", tc.answer)
				t.Logf(" got %v", tc.input)
			}
		})
	}
}

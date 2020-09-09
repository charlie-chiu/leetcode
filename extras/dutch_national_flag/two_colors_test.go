package dutch_national_flag

import (
	"reflect"
	"testing"
)

type TestCase struct {
	input, answer []int
}

func TestSort2Colors(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{0, 0}, []int{0, 0}},
		{[]int{1, 1}, []int{1, 1}},
		{[]int{1, 0, 0}, []int{0, 0, 1}},
		{[]int{1, 1, 1, 1, 0, 0, 0}, []int{0, 0, 0, 1, 1, 1, 1}},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			input := make([]int, len(tc.input))
			copy(input, tc.input)

			sort2Colors(tc.input)

			if !reflect.DeepEqual(tc.input, tc.answer) {
				t.Errorf("wrong answer on input %v", input)
				t.Logf("want %v", tc.answer)
				t.Logf(" got %v", tc.input)
			}
		})
	}
}

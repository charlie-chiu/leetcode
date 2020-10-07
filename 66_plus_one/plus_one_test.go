package plus_one

import (
	"reflect"
	"testing"
)

type TestCase struct {
	input, answer []int
}

func TestPlusOne(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{0}, []int{1}},
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{9}, []int{1, 0}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			input := make([]int, len(tc.input))
			copy(input, tc.input)

			got := plusOne(input)
			if !reflect.DeepEqual(got, tc.answer) {
				t.Errorf("wrong answer on input %v", tc.input)
				t.Logf("want %v got %v", tc.answer, got)
			}
		})
	}
}

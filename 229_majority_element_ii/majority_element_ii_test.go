package majority_element

import (
	"reflect"
	"testing"
)

type TestCase struct {
	input  []int
	answer []int
}

func TestMajorityElement(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{3, 2, 3}, []int{3}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 1, 5, 5, 5, 6}, []int{5}},
		{[]int{2, 2}, []int{2}},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			t.Run("", func(t *testing.T) {
				got := majorityElement(tc.input)
				if !reflect.DeepEqual(got, tc.answer) {
					t.Logf("wrong answer at input %v\n", tc.input)
					t.Logf("want %v got %v", tc.answer, got)
					t.Fail()
				}
			})
		})
	}
}

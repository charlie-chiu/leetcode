package find_disappeared_numbers

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums, answer []int
}

func TestFind(t *testing.T) {
	TestCases := []TestCase{
		{[]int{1, 1}, []int{2}},
		{[]int{3, 2, 1, 4}, []int{}},
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
	}

	for _, tc := range TestCases {
		got := findDisappearedNumbers(tc.nums)
		if !reflect.DeepEqual(got, tc.answer) {
			t.Errorf("wrong answer:\n  nums: %v\nanswer: %v\n   got: %v", tc.nums, tc.answer, got)
		}
	}

}

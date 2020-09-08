package two_sum_ii

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums   []int
	target int
	answer []int
}

func TestTwoSum(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
		{[]int{1, 2, 5, 8, 10}, 13, []int{3, 4}},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := twoSum(tc.nums, tc.target)

			if !reflect.DeepEqual(got, tc.answer) {
				t.Errorf("wrong answer:\n")
				t.Logf("numbers: %v", tc.nums)
				t.Logf(" target: %v", tc.target)
				t.Logf("   want: %v", tc.answer)
				t.Logf("    got: %v", got)
			}
		})
	}
}

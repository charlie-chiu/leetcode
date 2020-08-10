package two_sum

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums   []int
	target int
	answer []int
}

func TestSolution(t *testing.T) {

	testCases := []TestCase{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{11, 15, 4, 5}, 9, []int{2, 3}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}

	for _, tc := range testCases {
		got := solution(tc.nums, tc.target)

		if !reflect.DeepEqual(got, tc.answer) {
			t.Errorf("wrong answer: nums: %v, target: %v, answer: %v, got: %v", tc.nums, tc.target, tc.answer, got)
		}
	}
}

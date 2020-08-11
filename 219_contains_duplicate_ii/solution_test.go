package contains_duplicate

import (
	"testing"
)

type TestCase struct {
	nums []int
	k    int
	want bool
}

func TestSolution(t *testing.T) {
	testCases := []TestCase{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, false},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}

	for _, tc := range testCases {
		got := solution(tc.nums, tc.k)
		if got != tc.want {
			t.Errorf("wrong answer: nums: %v, k: %v, got: %v, want %v", tc.nums, tc.k, got, tc.want)
		}
	}
}

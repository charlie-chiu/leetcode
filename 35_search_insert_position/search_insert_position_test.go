package search_insert_position

import "testing"

type TestCase struct {
	nums   []int
	target int
	answer int
}

func TestSearchInsert(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1}, 0, 0},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := searchInsert(tc.nums, tc.target)
			if got != tc.answer {
				t.Errorf("wrong answer.")
				t.Logf("nums: %v, target: %v", tc.nums, tc.target)
				t.Logf("want: %d, got: %d", tc.answer, got)
			}
		})
	}
}

package minimum_path_sum

import "testing"

type TestCase struct {
	input  [][]int
	answer int
}

func TestMinimumPathSum(t *testing.T) {
	var TestCases = []TestCase{
		// 1-3-1-1-1
		{[][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		}, 7},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := minPathSum(tc.input)
			if got != tc.answer {
				t.Errorf("wrong answer, want %d got %d", tc.answer, got)
			}
		})
	}
}

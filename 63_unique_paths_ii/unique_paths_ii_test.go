package unique_paths_ii

import "testing"

type TestCase struct {
	grid   [][]int
	answer int
}

func TestUniquePaths(t *testing.T) {
	var TestCases = []TestCase{
		{[][]int{
			{0},
		}, 1},
		{[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}, 2},
		{[][]int{
			{0, 0, 0},
			{0, 1, 1},
			{0, 0, 0},
		}, 1},
		{[][]int{
			{0, 0, 0},
			{1, 1, 1},
			{0, 0, 0},
		}, 0},
		{[][]int{
			{0, 0, 1},
			{0, 0, 0},
			{1, 0, 0},
		}, 4},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := uniquePathsWithObstacles(tc.grid)
			if got != tc.answer {
				t.Errorf("wrong answer")
				t.Logf("wand %d got %d", tc.answer, got)
			}
		})
	}
}

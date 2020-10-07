package triangle

import "testing"

type TestCase struct {
	input  [][]int
	answer int
}

func TestTriangle(t *testing.T) {
	var TestCases = []TestCase{
		{[][]int{
			{1},
		}, 1},
		{[][]int{
			{1},
			{2, 5},
		}, 3},
		{[][]int{
			{2},
			{3, 4},
			{6, 5, 7},
			{4, 1, 8, 3},
		}, 11},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := minimumTotal(tc.input)
			if got != tc.answer {
				t.Errorf("wrong answer, want %d got %d", tc.answer, got)
			}
		})
	}
}

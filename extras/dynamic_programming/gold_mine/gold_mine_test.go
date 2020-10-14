package gold_mine

import "testing"

type TestCase struct {
	mine   [][]int
	answer int
}

func TestGetMaxGold(t *testing.T) {
	var TestCases = []TestCase{
		{[][]int{
			{0, 1},
			{0, 0},
		}, 1},
		{[][]int{
			{1, 3, 3},
			{2, 1, 4},
			{0, 6, 4},
		}, 12},
		{[][]int{
			{1, 3, 1, 5},
			{2, 2, 4, 1},
			{5, 0, 2, 3},
			{0, 6, 1, 2},
		}, 16,
		},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := getMaxGold(tc.mine)
			if got != tc.answer {
				t.Logf("wrong answer")
				t.Logf("want %d got %d", tc.answer, got)
				t.Fail()
			}
		})
	}
}

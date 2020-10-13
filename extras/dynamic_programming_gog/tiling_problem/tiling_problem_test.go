package tiling_problem

import "testing"

type TestCase struct {
	n, answer int
}

func TestCount(t *testing.T) {
	var TestCases = []TestCase{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := count(tc.n)
			if got != tc.answer {
				t.Logf("wrong answer")
				t.Logf("board: 2x%d, want %d got %d", tc.n, tc.answer, got)
				t.Fail()
			}
		})
	}
}

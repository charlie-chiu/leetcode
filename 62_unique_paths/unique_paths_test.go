package unique_paths

import "testing"

type TestCase struct {
	m, n   int
	answer int
}

func TestUniquePaths(t *testing.T) {
	var TestCases = []TestCase{
		{1, 1, 1},
		{1, 2, 1},
		{2, 1, 1},
		{3, 2, 3},
		{2, 3, 3},
		{3, 3, 6},
		{3, 7, 28},
		{51, 9, 1916797311},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := uniquePaths(tc.m, tc.n)
			if got != tc.answer {
				t.Errorf("wrong answer at m = %d, n = %d", tc.m, tc.n)
				t.Logf("wand %d got %d", tc.answer, got)
			}
		})
	}
}

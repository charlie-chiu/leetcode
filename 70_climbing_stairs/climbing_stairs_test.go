package climbing_stairs

import "testing"

type TestCase struct {
	top, answer int
}

func TestClimbingStairs(t *testing.T) {
	var TestCases = []TestCase{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{10, 89},
		{20, 10946},
		{45, 1836311903},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := climbStairs(tc.top)
			if got != tc.answer {
				t.Errorf("wrong answer")
				t.Errorf("top %d should have %d ways to climb, got %d", tc.top, tc.answer, got)
			}
		})
	}
}

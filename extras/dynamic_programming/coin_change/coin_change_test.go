package coin_change

import "testing"

type TestCase struct {
	values []int
	sum    int
	answer int
}

func TestCountWays(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{1, 2, 3}, 2, 2},
		{[]int{1, 2, 3}, 3, 3},
		{[]int{1, 2, 3}, 5, 5},
		{[]int{1, 3, 4, 5}, 7, 6},
		{[]int{1, 2, 3}, 4, 4},
		{[]int{2, 5, 3, 6}, 10, 5},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := countWays(tc.values, tc.sum)
			if got != tc.answer {
				t.Logf("wrong answer")
				t.Logf("values: %v, sum: %d", tc.values, tc.sum)
				t.Logf("want: %d got %d", tc.answer, got)
				t.Fail()
			}
		})
	}
}

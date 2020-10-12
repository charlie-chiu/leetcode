package best_time_to_buy_and_sell_stock

import "testing"

type TestCase struct {
	input  []int
	answer int
}

func TestMaxProfit(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{7}, 0},
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := maxProfit(tc.input)
			if got != tc.answer {
				t.Errorf("wrong answer on %v", tc.input)
				t.Logf("want %v got %v", tc.answer, got)
			}
		})
	}
}

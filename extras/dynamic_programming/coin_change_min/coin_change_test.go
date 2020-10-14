package coin_change_min

import "testing"

type TestCase struct {
	coins  []int
	amount int
	answer int
}

func TestCoinChange(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{1}, 1, 1},
		{[]int{1}, 2, 2},
		{[]int{1}, 0, 0},
		{[]int{25, 10, 5}, 30, 2},

		// greedy algorithm won't work on this case
		{[]int{9, 6, 5, 1}, 11, 2},
		{[]int{1, 3, 4, 5}, 7, 2},

		// leetcode : return -1 when amount of money can't be made up
		{[]int{2}, 3, -1},
		{[]int{3, 5}, 4, -1},

		//leetcode test case
		{[]int{186, 419, 83, 408}, 6249, 20},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := coinChange(tc.coins, tc.amount)
			if got != tc.answer {
				t.Logf("wrong answer")
				t.Logf("coins: %v, amount: %d", tc.coins, tc.amount)
				t.Logf("want %d got %d", tc.answer, got)
				t.Fail()
			}
		})
	}
}

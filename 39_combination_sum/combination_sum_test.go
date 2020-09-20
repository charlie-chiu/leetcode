package combination_sum

import (
	"reflect"
	"testing"
)

type TestCase struct {
	candidates []int
	target     int
	answer     [][]int
}

func TestCombinationSum(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{2, 3, 6, 7}, 7, [][]int{{7}, {2, 2, 3}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := combinationSum(tc.candidates, tc.target)

			if !reflect.DeepEqual(got, tc.answer) {
				t.Errorf("wrong answer")
				t.Logf("want %v", tc.answer)
				t.Logf(" got %v", got)
			}
		})
	}
}

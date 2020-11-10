package combination_sum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
			errMsg := fmt.Sprintf("target %d, candidates %v", tc.target, tc.candidates)
			assert.ElementsMatch(t, got, tc.answer, errMsg)
		})
	}
}

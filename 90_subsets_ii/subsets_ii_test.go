package subsets_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input  []int
	answer [][]int
}

func TestSubsets(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{0}, [][]int{{}, {0}}},
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {3}, {1, 2}, {2, 3}, {1, 3}, {1, 2, 3}}},

		{[]int{1, 2, 2}, [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}}},
		{[]int{2, 1, 2}, [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}}},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := subsetsWithDup(tc.input)
			assert.ElementsMatch(t, got, tc.answer)
		})
	}
}

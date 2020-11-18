package subsets

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

		// leetcode testcase, failed due to different order
		{[]int{9, 0, 3, 5, 7}, [][]int{{}, {9}, {0}, {0, 9}, {3}, {3, 9}, {0, 3}, {0, 3, 9}, {5}, {5, 9}, {0, 5}, {0, 5, 9}, {3, 5}, {3, 5, 9}, {0, 3, 5}, {0, 3, 5, 9}, {7}, {7, 9}, {0, 7}, {0, 7, 9}, {3, 7}, {3, 7, 9}, {0, 3, 7}, {0, 3, 7, 9}, {5, 7}, {5, 7, 9}, {0, 5, 7}, {0, 5, 7, 9}, {3, 5, 7}, {3, 5, 7, 9}, {0, 3, 5, 7}, {0, 3, 5, 7, 9}}},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := subsets(tc.input)
			assert.ElementsMatch(t, got, tc.answer)
		})
	}
}

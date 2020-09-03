package path_sum

import (
	"testing"

	. "leetcode"
)

type TestCase struct {
	tree   *TreeNode
	sum    int
	answer bool
}

func TestPathSum(t *testing.T) {
	var TestCases = []TestCase{
		{nil, 87, false},
		{nil, 0, false},
		{NewTree(1), 1, true},
		{NewTree(1, 2), 1, false},
		{NewTree(1, 2, 3), 6, false},
		{NewTree(1, 2, 3, 5), 8, true},
		{NewTree(5, 4, 8, 11, 13, 4, 7, 2, nil, nil, 1), 22, true},
	}

	for _, tc := range TestCases {
		got := hasPathSum(tc.tree, tc.sum)
		if got != tc.answer {
			t.Errorf("wrong answer.")
			t.Logf("finding sum %d in tree %v", tc.sum, tc.tree)
			t.Logf("want: %v", tc.answer)
			t.Logf(" got: %v", got)
		}
	}
}

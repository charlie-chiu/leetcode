package maximum_depth

import (
	"testing"

	. "leetcode"
)

type TestCase struct {
	tree   *TreeNode
	output int
}

func TestMaximumDepth(t *testing.T) {
	var TestCases = []TestCase{
		{NewTree(1), 1},
		{NewTree(nil), 0},
		{NewTree(1, 2), 2},
		{NewTree(3, 9, 20, nil, nil, 15, 7), 2},
	}

	for _, testCase := range TestCases {
		got := minDepth(testCase.tree)

		if got != testCase.output {
			t.Errorf("assert min depth of %v, is %d, got %d", testCase.tree, testCase.output, got)
		}
	}
}

package balanced_binary_tree

import (
	"reflect"
	"testing"

	. "leetcode"
)

type TestCase struct {
	tree   *TreeNode
	output bool
}

func TestIsBalanced(t *testing.T) {
	var TestCases = []TestCase{
		//{NewTree(1), true},
		//{NewTree(1, nil, 2), true},
		{NewTree(1, nil, 2, 3), false},
		{NewTree(3, 9, 20, nil, nil, 15, 7), true},
		{NewTree(1, 2, 2, 3, 3, nil, nil, 4, 4), false},
		{NewTree(1, 2, 2, 3, nil, nil, 3, 4, nil, nil, 4), false},
	}

	for _, tc := range TestCases {
		got := isBalanced(tc.tree)

		if !reflect.DeepEqual(got, tc.output) {
			t.Errorf("wrong output at tree %v, got %v", tc.tree, got)
		}
	}
}

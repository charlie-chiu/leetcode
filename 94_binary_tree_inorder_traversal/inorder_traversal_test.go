package binary_tree_inorder_traversal

import (
	"reflect"

	"testing"

	. "leetcode"
)

type TestCase struct {
	input  *TreeNode
	answer []int
}

func TestInorderTraversal(t *testing.T) {
	var TestCases = []TestCase{
		{nil, []int{}},
		{NewTree(1, nil, 2, 3), []int{1, 3, 2}},
		{NewTree(1, 4, 3, 5, 6, 2), []int{5, 4, 6, 1, 2, 3}},
	}

	for _, tc := range TestCases {
		got := inorderTraversal(tc.input)

		if !reflect.DeepEqual(got, tc.answer) {
			t.Errorf("wrong answer")
			t.Logf("tree: %v", tc.input)
			t.Logf("want: %v", tc.answer)
			t.Logf(" got: %v", got)
		}
	}
}

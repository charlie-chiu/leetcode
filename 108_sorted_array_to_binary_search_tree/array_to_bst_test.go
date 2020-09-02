package array_to_bst

import (
	"reflect"
	"testing"

	. "leetcode"
)

type TestCase struct {
	input  []int
	answer *TreeNode
}

func TestArrayToBST(t *testing.T) {
	// may have multi possible answer, should write a BST validator here
	var TestCases = []TestCase{
		{[]int{}, nil},
		{[]int{-10, -3, 0, 5, 9}, NewTree(0, -3, 9, -10, nil, 5)},
		{[]int{2, 5, 7, 9}, NewTree(7, 5, 9, 2)},
		{[]int{0, 1, 2, 3, 4, 5}, NewTree(3, 1, 5, 0, 2, 4)},
	}

	for _, tc := range TestCases {
		got := sortedArrayToBST(tc.input)

		if !reflect.DeepEqual(got, tc.answer) {
			t.Errorf("wrong answer.")
			t.Logf("   input: %v", tc.input)
			t.Logf("expected: %v", tc.answer)
			t.Logf("     got: %v", got)
		}
	}
}

//
//func isBST(root *TreeNode) bool {
//	//base case
//	if root == nil {
//		return true
//	}
//
//	if root.Left == nil && root.Right == nil {
//		return true
//	}
//
//}

package level_order_traversal

import (
	"reflect"
	"testing"

	. "leetcode"
)

type TestCase struct {
	input  *TreeNode
	answer [][]int
}

func TestLevelOrderTraversal(t *testing.T) {
	var TestCases = []TestCase{
		{NewTree(nil), [][]int{}},
		{NewTree(1), [][]int{{1}}},
		{NewTree(1, 2, 3), [][]int{{1}, {2, 3}}},
		{NewTree(3, 9, 20, 15, 7), [][]int{{3}, {9, 20}, {15, 7}}},
	}

	for _, tc := range TestCases {
		got := levelOrder(tc.input)
		if !reflect.DeepEqual(got, tc.answer) {
			t.Errorf("wrong answer on input %v", tc.input)
			t.Logf("expected: %v", tc.answer)
			t.Logf("     got: %v", got)
		}
	}
}

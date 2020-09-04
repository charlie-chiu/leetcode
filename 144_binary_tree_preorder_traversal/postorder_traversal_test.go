package preorder_traversal

import (
	"reflect"
	"testing"

	. "leetcode"
)

type TestCase struct {
	input  *TreeNode
	answer []int
}

func TestPostorderTraversal(t *testing.T) {
	var TestCases = []TestCase{
		{nil, []int{}},
		{NewTree(1), []int{1}},
		{NewTree(1, 2), []int{1, 2}},
		{NewTree(1, nil, 2), []int{1, 2}},
		{NewTree(1, nil, 2, 3), []int{1, 2, 3}},
		{NewTree(1, 4, 3, 5, 6, 2), []int{1, 4, 5, 6, 3, 2}},
	}

	for _, tc := range TestCases {
		got := preorderTraversal(tc.input)

		if !reflect.DeepEqual(got, tc.answer) {
			t.Errorf("wrong answer")
			t.Logf("input: %v", tc.input)
			t.Logf(" want: %v", tc.answer)
			t.Logf("  got: %v", got)
		}
	}
}

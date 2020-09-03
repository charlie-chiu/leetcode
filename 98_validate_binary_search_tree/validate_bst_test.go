package validate_binary_search_tree

import (
	"reflect"
	"testing"

	. "leetcode"
)

type TestCase struct {
	input  *TreeNode
	answer bool
}

func TestIsValidBST(t *testing.T) {
	var TestCases = []TestCase{
		// is nil a tree?
		{NewTree(nil), true},

		{NewTree(0, -1), true},
		{NewTree(1, 1), false},
		{NewTree(2), true},
		{NewTree(2, 1, 3), true},
		{NewTree(5, 1, 4, nil, nil, 3, 6), false},

		{NewTree(10, 5, 15, nil, nil, 6, 20), false},

		//balanced?
		{NewTree(5, 1, nil, 0), false},
	}

	for _, tc := range TestCases {
		got := isValidBST(tc.input)
		if !reflect.DeepEqual(got, tc.answer) {
			t.Errorf("wrong answer")
			t.Logf("   input %v", tc.input)
			t.Logf("expected %v", tc.answer)
			t.Logf("     got %v", got)
		}

	}
}

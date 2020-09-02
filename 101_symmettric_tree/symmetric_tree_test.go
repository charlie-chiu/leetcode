package symmettric_tree

import (
	"reflect"
	"testing"

	. "leetcode"
)

type TestCase struct {
	tree *TreeNode
	want bool
}

func TestIsSymmetric(t *testing.T) {
	var TestCases = []TestCase{
		{NewTree(), true},
		{NewTree(1), true},
		{NewTree(1, 2, 3), false},
		{NewTree(1, 2, 2), true},
		{NewTree(1, 2, 2, 3, 4, 4, 3), true},
		{NewTree(1, 2, 2, nil, 3, nil, 3), false},
	}

	for _, testCase := range TestCases {
		got := isSymmetric(testCase.tree)

		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("assert failed on tree %v, got %v", testCase.tree, got)
		}
	}
}

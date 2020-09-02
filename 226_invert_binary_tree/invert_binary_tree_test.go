package invert_binary_tree

import (
	"reflect"
	"testing"

	. "leetcode"
)

type TestCase struct {
	input, output *TreeNode
}

func TestInvertTree(t *testing.T) {
	var TestCases = []TestCase{
		{NewTree(1), NewTree(1)},
		{NewTree(1, 2, 3), NewTree(1, 3, 2)},
		{NewTree(4, 2, 7, 1, 3, 6, 9), NewTree(4, 7, 2, 9, 6, 3, 1)},
	}

	for _, tc := range TestCases {
		got := invertTree(tc.input)

		if !reflect.DeepEqual(got, tc.output) {
			t.Fail()
			t.Logf("wrong answer on tree %v", tc.input)
			t.Logf("want: %v", tc.output)
			t.Logf(" got: %v", got)
		}
	}
}

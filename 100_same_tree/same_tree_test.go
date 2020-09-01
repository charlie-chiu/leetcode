package same_tree

import (
	"testing"

	. "leetcode"
)

type TestCase struct {
	p, q   *TreeNode
	output bool
}

func TestIsSame(t *testing.T) {
	var TestCases = []TestCase{
		{NewTree(nil), NewTree(nil), true},
		{NewTree(1, 1, 2), NewTree(1, 1, 2), true},
		{NewTree(1, 2), NewTree(1, nil, 2), false},
		{NewTree(1, 2), NewTree(2, 1), false},
	}

	for _, testCase := range TestCases {
		got := isSameTree(testCase.p, testCase.q)

		if got != testCase.output {
			t.Errorf("wrong answer on tree %q and %q", testCase.p, testCase.q)
			t.Logf("want %v got %v", testCase.output, got)
		}
	}
}

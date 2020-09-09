package linked_list_cycle_ii

import (
	"reflect"
	"testing"

	. "leetcode"
)

type TestCase struct {
	input, answer *ListNode
}

func TestMiddleNode(t *testing.T) {
	var TestCases = []TestCase{
		{newCyclicList([]int{}, -1), nil},
		{newCyclicList([]int{1}, -1), nil},
		{newCyclicList([]int{1, 2}, -1), nil},
		{newCyclicList([]int{3, 2, 0, -4}, 1), newCyclicList([]int{2, 0, -4}, 0)},
		{newCyclicList([]int{1, 2}, 0), newCyclicList([]int{1, 2}, 0)},
	}

	for _, testCase := range TestCases {
		got := detectCycle(testCase.input)

		if !reflect.DeepEqual(testCase.answer, got) {
			t.Errorf("wrong answer")
			t.Logf("input")
			DumpCyclicLst(testCase.input)
			t.Logf("want")
			DumpCyclicLst(testCase.answer)
			t.Logf(" got")
			DumpCyclicLst(got)
		}
	}
}

func newCyclicList(n []int, pos int) *ListNode {
	if len(n) < 1 {
		return nil
	}
	nodes := make([]*ListNode, len(n))

	for i := 0; i < len(n); i++ {
		nodes[i] = &ListNode{
			Val:  n[i],
			Next: nil,
		}
	}
	for i := 0; i < len(n); i++ {
		if i == len(n)-1 {
			if pos > -1 {
				nodes[i].Next = nodes[pos]
			}
		} else {
			nodes[i].Next = nodes[i+1]
		}
	}

	return nodes[0]
}

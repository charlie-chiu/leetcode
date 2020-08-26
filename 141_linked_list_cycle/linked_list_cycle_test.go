package linked_list_cycle

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nodeValues *ListNode
	want       bool
}

func TestMiddleNode(t *testing.T) {
	var TestCases = []TestCase{
		{newCyclicList([]int{}, -1), false},
		{newCyclicList([]int{3, 2, 0, -4}, 1), true},
		{newCyclicList([]int{1, 2}, 0), true},
		{newCyclicList([]int{1}, -1), false},
	}

	for i, testCase := range TestCases {
		got := hasCycle(testCase.nodeValues)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("wrong output at case %d", i)
			t.Logf("want %v", testCase.want)
			t.Logf(" got %v", got)
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

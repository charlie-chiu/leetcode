package merge_two_sorted_lists

import (
	"reflect"
	"testing"

	. "leetcode"
)

type TestCase struct {
	L1, L2, Answer *ListNode
}

func TestMergeTwoLists(t *testing.T) {
	var TestCases = []TestCase{
		{nil, nil, nil},
		{NewLinkedList(1), NewLinkedList(1), NewLinkedList(1, 1)},
		{NewLinkedList(1, 2, 3), nil, NewLinkedList(1, 2, 3)},
		{NewLinkedList(1, 2, 4), NewLinkedList(1, 3, 4), NewLinkedList(1, 1, 2, 3, 4, 4)},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := mergeTwoLists(tc.L1, tc.L2)
			if !reflect.DeepEqual(got, tc.Answer) {
				t.Errorf("Wrong answer")
				t.Logf("merge %v and %v", tc.L1, tc.L2)
				t.Logf("want %v", tc.Answer)
				t.Logf(" got %v", got)
			}
		})
	}
}

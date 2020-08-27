package add_two_numbers

import (
	"reflect"
	"testing"

	. "leetcode"
)

type TestCase struct {
	L1, L2, output *ListNode
}

func TestAddTwoNumbers(t *testing.T) {
	var TestCases = []TestCase{
		{NewLinkedList(1, 8), NewLinkedList(0), NewLinkedList(1, 8)},
		{NewLinkedList(5), NewLinkedList(5), NewLinkedList(0, 1)},
		{NewLinkedList(2, 4, 3), NewLinkedList(5, 6, 4), NewLinkedList(7, 0, 8)},
	}

	for _, tc := range TestCases {
		got := addTwoNumbers(tc.L1, tc.L2)

		if !reflect.DeepEqual(got, tc.output) {
			t.Errorf("wrong ansewr")
			t.Logf("  L1:%v", tc.L1)
			t.Logf("  L2:%v", tc.L2)
			t.Logf("want:%v", tc.output)
			t.Logf(" got:%v", got)
		}
	}
}

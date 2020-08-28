package add_two_numbers

import (
	"fmt"
	"reflect"
	"testing"

	. "leetcode"
)

type TestCase struct {
	l1, l2, sum *ListNode
}

func TestAddTwoNumbers(t *testing.T) {
	var TestCases = []TestCase{
		{NewLinkedList(4), NewLinkedList(6), NewLinkedList(1, 0)},
		{NewLinkedList(7, 2, 4, 3), NewLinkedList(5, 6, 4), NewLinkedList(7, 8, 0, 7)},
	}

	for i, tc := range TestCases {
		testName := fmt.Sprintf("testcase %d", i)
		t.Run(testName, func(t *testing.T) {
			got := addTwoNumbers(tc.l1, tc.l2)
			if !reflect.DeepEqual(got, tc.sum) {
				t.Errorf("wrong ansewr")
				t.Logf("  L1:%v", tc.l1)
				t.Logf("  L2:%v", tc.l2)
				t.Logf("want:%v", tc.sum)
				t.Logf(" got:%v", got)
			}
		})
	}
}

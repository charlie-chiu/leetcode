package delete_node

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	t.Run(fmt.Sprintf("test 4,5,1,9, node 5"), func(t *testing.T) {
		list := &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
		}

		deleteNode(list.Next)

		want := &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		}
		if !reflect.DeepEqual(list, want) {
			t.Errorf("wrong output")
			t.Logf("want %v", want)
			t.Logf(" got %v", list)
		}
	})
}

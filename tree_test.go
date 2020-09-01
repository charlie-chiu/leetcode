package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var tree *TreeNode = nil
		want := "<nil>"
		got := fmt.Sprintf("%s", tree)

		if got != want {
			t.Errorf("wrong output,want %q,  got %q", want, got)
		}

	})

	t.Run("basic tree", func(t *testing.T) {
		var tree = &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 3},
			},
		}
		want := "1,<nil>,2,3"
		got := fmt.Sprintf("%s", tree)

		if got != want {
			t.Errorf("wrong output,want %q,  got %q", want, got)
		}

	})
}

func TestNewTree(t *testing.T) {
	t.Run("nil return nil", func(t *testing.T) {
		var got = NewTree(nil)
		var want *TreeNode = nil

		assertTreeSame(t, want, got)
	})

	t.Run("[1, nil, 2, 3]", func(t *testing.T) {
		var got = NewTree(1, nil, 2, 3)
		var want = &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 3},
			},
		}

		assertTreeSame(t, want, got)
	})
}

func assertTreeSame(t *testing.T, want, got *TreeNode) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("tree not equal")
		t.Logf("want: %v", want)
		t.Logf(" got: %v", got)
	}
}

package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(values ...interface{}) *TreeNode {
	return nil
}

// level order traversal
func (t *TreeNode) String() string {
	if t == nil {
		return fmt.Sprint(nil)
	}

	var result []string
	queue := []*TreeNode{t}
	for len(queue) > 0 {
		// dequeue
		curr := queue[0]
		queue = queue[1:]

		if curr != nil {
			result = append(result, strconv.Itoa(curr.Val))
			queue = append(queue, curr.Left)
			queue = append(queue, curr.Right)
		} else {
			result = append(result, fmt.Sprint(nil))
		}

	}
	for result[len(result)-1] == fmt.Sprint(nil) {
		result = result[:len(result)-1]
	}

	return strings.Join(result, ",")
}

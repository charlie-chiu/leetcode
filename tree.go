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
	if len(values) < 1 || values[0] == nil {
		return nil
	}
	if len(values) == 1 {
		return &TreeNode{Val: values[0].(int)}
	}

	var treeQueue []*TreeNode
	for _, value := range values {
		treeQueue = append(treeQueue, NewTree(value))
	}

	var root *TreeNode
	var parentQueue = []**TreeNode{}
	for len(treeQueue) > 0 {
		//dequeue
		curr := treeQueue[0]
		treeQueue = treeQueue[1:]

		//set root
		if len(parentQueue) == 0 {
			root = curr
		} else {
			*parentQueue[0] = *&curr
			parentQueue = parentQueue[1:]
		}

		if curr != nil {
			parentQueue = append(parentQueue, &curr.Left)
			parentQueue = append(parentQueue, &curr.Right)
		}
	}

	return root
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

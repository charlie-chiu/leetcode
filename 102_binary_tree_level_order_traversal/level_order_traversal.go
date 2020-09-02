package level_order_traversal

import (
	. "leetcode"
)

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	var nextLevelQueue []*TreeNode
	var currLevelQueue []*TreeNode = []*TreeNode{root}

	var currNode *TreeNode
	var currValue []int
	for len(currLevelQueue) > 0 {
		//dequeue first
		currNode = currLevelQueue[0]
		currLevelQueue = currLevelQueue[1:]

		currValue = append(currValue, currNode.Val)

		if currNode.Left != nil {
			nextLevelQueue = append(nextLevelQueue, currNode.Left)
		}

		if currNode.Right != nil {
			nextLevelQueue = append(nextLevelQueue, currNode.Right)
		}

		if len(currLevelQueue) == 0 {
			result = append(result, currValue)
			currValue = []int{}

			currLevelQueue = nextLevelQueue
			nextLevelQueue = []*TreeNode{}
		}
	}

	return result
}

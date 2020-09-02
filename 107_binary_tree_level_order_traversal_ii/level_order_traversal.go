package level_order_traversal_ii

import (
	. "leetcode"
)

// 12ms , 6.6MB poor performance
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	var nextLevelQueue []*TreeNode
	var currLevelQueue []*TreeNode = []*TreeNode{root}

	var currNode *TreeNode
	var currValues []int
	for len(currLevelQueue) > 0 {
		//dequeue first
		currNode = currLevelQueue[0]
		currLevelQueue = currLevelQueue[1:]

		currValues = append(currValues, currNode.Val)

		if currNode.Left != nil {
			nextLevelQueue = append(nextLevelQueue, currNode.Left)
		}

		if currNode.Right != nil {
			nextLevelQueue = append(nextLevelQueue, currNode.Right)
		}

		if len(currLevelQueue) == 0 {
			result = append([][]int{currValues}, result...)
			currValues = []int{}

			currLevelQueue = nextLevelQueue
			nextLevelQueue = []*TreeNode{}
		}
	}

	return result
}

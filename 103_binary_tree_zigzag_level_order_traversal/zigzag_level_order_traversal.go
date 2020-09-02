package zigzag_level_order_traversal

import (
	. "leetcode"
)

// 0ms, 2.8MB
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	var nextLevelQueue []*TreeNode
	var currLevelQueue []*TreeNode = []*TreeNode{root}

	var currNode *TreeNode
	var currValues []int

	var fromLeft bool = true
	for len(currLevelQueue) > 0 {
		//dequeue first
		currNode = currLevelQueue[0]
		currLevelQueue = currLevelQueue[1:]

		if fromLeft {
			// use as queue append from last
			currValues = append(currValues, currNode.Val)
		} else {
			// use as stack ,push and pop from front
			currValues = append([]int{currNode.Val}, currValues...)
		}

		if currNode.Left != nil {
			nextLevelQueue = append(nextLevelQueue, currNode.Left)
		}
		if currNode.Right != nil {
			nextLevelQueue = append(nextLevelQueue, currNode.Right)
		}

		if len(currLevelQueue) == 0 {
			result = append(result, currValues)
			currValues = []int{}

			currLevelQueue = nextLevelQueue
			nextLevelQueue = []*TreeNode{}

			fromLeft = !fromLeft
		}
	}

	return result
}

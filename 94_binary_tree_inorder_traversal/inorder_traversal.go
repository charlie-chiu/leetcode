package binary_tree_inorder_traversal

import (
	. "leetcode"
)

// 0ms, 2MB
func inorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var node = root
	var answer []int = make([]int, 0)

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		answer = append(answer, node.Val)

		node = node.Right
	}

	return answer
}

// first try with recursion
// 0ms, 2MB
//func inorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return []int{}
//	}
//
//	return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
//}

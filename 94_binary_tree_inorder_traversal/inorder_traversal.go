package binary_tree_inorder_traversal

import (
	. "leetcode"
)

// 0ms, 2MB
func inorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var result = make([]int, 0)

	var curr = root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		//pop
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, curr.Val)

		curr = curr.Right
	}

	return result
}

// first try with recursion
// 0ms, 2MB
//func inorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return []int{}
//	}
//
//	left := inorderTraversal(root.Left)
//	right := inorderTraversal(root.Right)
//	return append(append(left, root.Val), right...)
//}

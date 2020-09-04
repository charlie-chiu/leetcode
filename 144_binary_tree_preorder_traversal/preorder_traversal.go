package preorder_traversal

import . "leetcode"

// 0ms, 2.1MB
func preorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var result = make([]int, 0)

	var node = root
	for node != nil || len(stack) > 0 {
		for node != nil {
			result = append(result, node.Val)
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node = node.Right
	}

	return result
}

// 0ms, 2.1MB
func recursion(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	left := recursion(root.Left)
	right := recursion(root.Right)
	return append(append([]int{root.Val}, left...), right...)
}

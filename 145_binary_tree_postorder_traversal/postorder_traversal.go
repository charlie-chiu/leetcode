package postorder_traversal

import . "leetcode"

// 0ms, 2.1MB
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	left := postorderTraversal(root.Left)
	right := postorderTraversal(root.Right)
	return append(append(left, right...), root.Val)
}

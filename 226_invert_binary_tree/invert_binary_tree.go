package invert_binary_tree

import . "leetcode"

// 0ms, 2.2MB, not good at memory usage
func invertTree(root *TreeNode) *TreeNode {
	// base case
	if root == nil {
		return nil
	}

	return &TreeNode{
		Val:   root.Val,
		Left:  invertTree(root.Right),
		Right: invertTree(root.Left),
	}
}

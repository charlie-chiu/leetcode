package balanced_binary_tree

import . "leetcode"

// got 8ms, 86.48% and 5.7MB, 87.30% on sep 2 2020, serious?
func isBalanced(root *TreeNode) bool {
	// brute force with recursion
	// base case
	if root == nil {
		return true
	}

	diff := maxDepth(root.Left) - maxDepth(root.Right)
	currIsBalanced := diff >= -1 && diff <= 1

	return currIsBalanced && isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

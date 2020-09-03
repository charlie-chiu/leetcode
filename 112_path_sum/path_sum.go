package path_sum

import . "leetcode"

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val == sum {
		return true
	}

	if hasPathSum(root.Left, sum-root.Val) {
		return true
	}
	if hasPathSum(root.Right, sum-root.Val) {
		return true
	}

	return false
}

package symmettric_tree

import (
	. "leetcode"
)

// accepted on first try
// 0ms, 3MB
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirrored(root.Left, root.Right)
}

func isMirrored(p, q *TreeNode) bool {
	//base cases
	if p == nil && q == nil {
		//both nil
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isMirrored(p.Left, q.Right) && isMirrored(p.Right, q.Left)
}

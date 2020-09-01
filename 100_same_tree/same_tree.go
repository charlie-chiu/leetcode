package same_tree

import (
	. "leetcode"
)

//Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

// 0ms, 2.1MB
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// compare p & q
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	// compare children
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

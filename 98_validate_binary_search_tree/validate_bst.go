package validate_binary_search_tree

import (
	"math"

	. "leetcode"
)

// leetcode solution
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lowerLimit, upperLimit int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lowerLimit {
		return false
	}
	if root.Val >= upperLimit {
		return false
	}
	if !helper(root.Right, root.Val, upperLimit) {
		return false
	}
	if !helper(root.Left, lowerLimit, root.Val) {
		return false
	}

	return true
}

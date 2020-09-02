package array_to_bst

import (
	. "leetcode"
)

func sortedArrayToBST(nums []int) *TreeNode {
	//base case
	length := len(nums)
	if length == 0 {
		return nil
	}

	if length == 1 {
		return &TreeNode{Val: nums[0]}
	}

	middle := length / 2
	return &TreeNode{
		Val:   nums[middle],
		Left:  sortedArrayToBST(nums[:middle]),
		Right: sortedArrayToBST(nums[middle+1:]),
	}
}

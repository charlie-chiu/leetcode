package postorder_traversal

import . "leetcode"

// 0ms, 2.1MB
func postorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var result = make([]int, 0)

	var curr = root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr.Right)
			stack = append(stack, curr)
			curr = curr.Left
		}

		//pop
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//peek
		if len(stack) > 0 && curr.Right == stack[len(stack)-1] {
			// pop
			stack = stack[:len(stack)-1]
			stack = append(stack, curr)
			curr = curr.Right
		} else {
			result = append(result, curr.Val)
			curr = nil
		}
	}

	return result
}

// 0ms, 2.1MB
//func postorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return []int{}
//	}
//
//	left := postorderTraversal(root.Left)
//	right := postorderTraversal(root.Right)
//	return append(append(left, right...), root.Val)
//}

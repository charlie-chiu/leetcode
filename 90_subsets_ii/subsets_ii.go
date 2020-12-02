package subsets_ii

import "sort"

// first, recursive approach
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	return helper([]int{}, nums)
}

func helper(path []int, nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{path}
	}

	firstNum := nums[0]
	nums = nums[1:]
	var choosePath = make([]int, len(path))
	copy(choosePath, path)

	// with duplicate
	if len(path) > 0 && firstNum == path[len(path)-1] {
		return helper(append(choosePath, firstNum), nums)
	}

	// without duplicate
	return append(helper(append(choosePath, firstNum), nums), helper(path, nums)...)
}

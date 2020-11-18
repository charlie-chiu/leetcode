package subsets

// iterative approach
func subsets(nums []int) [][]int {
	var subsets = [][]int{{}}

	for _, num := range nums {
		for _, subset := range subsets {
			newSubset := make([]int, len(subset), len(subset)+1)
			copy(newSubset, subset)
			newSubset = append(newSubset, num)
			subsets = append(subsets, newSubset)
		}
	}

	return subsets
}

// recurse approach
//func subsets(nums []int) [][]int {
//	return subsetsHelper([]int{}, nums)
//}
//
//func subsetsHelper(path []int, nums []int) [][]int {
//	//base case, no more num to choose
//	if len(nums) == 0 {
//		return [][]int{path}
//	}
//
//	firstNum := nums[0]
//	nums = nums[1:]
//	var choosePath = make([]int, len(path))
//	copy(choosePath, path)
//	// choose first number                             //not choose first number
//	return append(subsetsHelper(append(choosePath, firstNum), nums), subsetsHelper(path, nums)...)
//}

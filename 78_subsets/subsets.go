package subsets

func subsets(nums []int) [][]int {
	return subsetsHelper([]int{}, nums)
}

func subsetsHelper(path []int, nums []int) [][]int {
	//base case, no more num to choose
	if len(nums) == 0 {
		return [][]int{path}
	}

	firstNum := nums[0]
	nums = nums[1:]
	var choosePath = make([]int, len(path))
	copy(choosePath, path)
	// choose first number                             //not choose first number
	return append(subsetsHelper(append(choosePath, firstNum), nums), subsetsHelper(path, nums)...)
}

// first try myself
//func subsets(nums []int) [][]int {
//	// base case
//	if len(nums) == 1 {
//		return [][]int{{}, {nums[0]}}
//	}
//
//	num := nums[len(nums)-1]
//	subsetsWithoutNum := subsets(nums[:len(nums)-1])
//	var subsetWithNum = make([][]int, len(subsetsWithoutNum))
//	copy(subsetWithNum, subsetsWithoutNum)
//
//	for i, subset := range subsetWithNum {
//		temp := make([]int, len(subset))
//		copy(temp, subset)
//		subsetWithNum[i] = append(temp, num)
//	}
//
//	return append(subsetsWithoutNum, subsetWithNum...)
//}

package subsets

func subsets(nums []int) [][]int {
	// base case
	if len(nums) == 1 {
		return [][]int{{}, {nums[0]}}
	}

	num := nums[len(nums)-1]
	subsetsWithoutNum := subsets(nums[:len(nums)-1])
	var subsetWithNum = make([][]int, len(subsetsWithoutNum))
	copy(subsetWithNum, subsetsWithoutNum)

	for i, subset := range subsetWithNum {
		temp := make([]int, len(subset))
		copy(temp, subset)
		subsetWithNum[i] = append(temp, num)
	}

	return append(subsetsWithoutNum, subsetWithNum...)
}

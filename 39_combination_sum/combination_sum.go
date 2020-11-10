package combination_sum

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int

	helper(candidates, target, []int{}, &result)

	return result
}

func helper(candidates []int, target int, path []int, result *[][]int) {
	if target < 0 {
		return
	}

	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)

		*result = append(*result, temp)
		return
	}

	for i, candidate := range candidates {
		helper(candidates[i:], target-candidate, append(path, candidate), result)
	}
}

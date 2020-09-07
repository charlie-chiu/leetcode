package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var uniqueIndex int

	for R := 0; R < len(nums); R++ {
		if nums[R] > nums[uniqueIndex] {
			uniqueIndex++
			nums[uniqueIndex] = nums[R]
		}
	}

	return uniqueIndex + 1
}

package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	const maxAppear = 2

	var slow int = 0
	var count int = 1

	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] > nums[slow] {
			// replace with bigger number
			slow++
			nums[slow] = nums[fast]
			count = 1
		} else if nums[fast] == nums[slow] && count < maxAppear {
			// num appeared
			slow++
			nums[slow] = nums[fast]
			count++
		}
	}

	return slow + 1
}

package majority_element

func majorityElement(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}

	counter := make(map[int]int, 0)

	for _, num := range nums {
		counter[num]++

		if counter[num] > length/2 {
			return num
		}
	}

	return 0
}

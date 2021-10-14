package majority_element

// first try
// got 16ms and 5.2 MB, poor result
func majorityElement(nums []int) []int {
	numsLen := len(nums)

	if numsLen == 1 {
		return nums
	}

	if numsLen == 2 {
		if nums[0] == nums[1] {
			return []int{nums[0]}
		} else {
			return nums
		}
	}

	counter := make(map[int]int)
	answer := make([]int, 0, 2)
	for _, num := range nums {
		counter[num]++
	}

	for num, count := range counter {
		if count > numsLen/3 {
			answer = append(answer, num)
		}
	}

	return answer
}

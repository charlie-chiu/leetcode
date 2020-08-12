package move_zeros

var moveZeroes func([]int) = inPlace

func inPlace(nums []int) {
	var remove = func(from []int, at int) []int {
		return append(from[:at], from[at+1:]...)
	}

	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[index] == 0 {
			nums = remove(nums, index)
			nums = append(nums, 0)
		} else {
			index++
		}
	}
}

func additionalSpace(nums []int) {
	var zeroCount int = 0
	var newArray []int

	for _, num := range nums {
		if num == 0 {
			zeroCount++
		} else {
			newArray = append(newArray, num)
		}
	}

	// add zeros
	for i := 0; i < zeroCount; i++ {
		newArray = append(newArray, 0)
	}

	copy(nums, newArray)
}

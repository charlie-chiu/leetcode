package two_sum_ii

func twoSum(numbers []int, target int) []int {
	// first try
	// 4ms(94.98%), 3MB
	var L, R = 0, len(numbers) - 1

	for numbers[L]+numbers[R] != target {
		sum := numbers[L] + numbers[R]
		if sum > target {
			R--
		}
		if sum < target {
			L++
		}
	}

	return []int{L + 1, R + 1}
}

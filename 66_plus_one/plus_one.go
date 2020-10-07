package plus_one

func plusOne(digits []int) []int {
	i := len(digits) - 1

	// consider plus one is carry
	var carry = 1
	var sum int

	for carry != 0 && i >= 0 {
		sum = (digits[i] + carry) % 10
		carry = (digits[i] + carry) / 10
		digits[i] = sum
		i--
	}

	if carry > 0 {
		return append([]int{carry}, digits...)
	}

	return digits
}

package power_of_two

// 4ms / 2.2MB
func isPowerOfTwo(n int) bool {
	powerOfTwo := 1
	for powerOfTwo < n {
		powerOfTwo = powerOfTwo << 1
	}

	return powerOfTwo == n
}

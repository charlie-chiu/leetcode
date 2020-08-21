package power_of_two

import "math/bits"

var solution = lc2

// from leetcode discussion
// 0ms / 2.2MB
func lc2(n int) bool {
	return bits.OnesCount(uint(n)) == 1
}

// from leetcode discussion
// 4ms / 2.2MB
func lc1(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

// first try
// 4ms / 2.2MB
func isPowerOfTwo(n int) bool {
	powerOfTwo := 1
	for powerOfTwo < n {
		powerOfTwo = powerOfTwo << 1
	}

	return powerOfTwo == n
}

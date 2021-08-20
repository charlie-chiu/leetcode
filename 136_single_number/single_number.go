package single_number

// just using XOR operator
// is there any others ways between xor and brute force?
func singleNumber(nums []int) int {
	var result int

	for _, num := range nums {
		result ^= num
	}

	return result
}

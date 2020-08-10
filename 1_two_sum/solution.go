package two_sum

var solution = firstTry

// first try: brute force
// time complexity: O(n!)
func firstTry(nums []int, target int) []int {
	for firstIndex, num := range nums {
		for i2, num2 := range nums[firstIndex+1:] {
			if num+num2 == target {
				return []int{firstIndex, i2 + firstIndex + 1}
			}
		}
	}

	// each input would have exactly one solution
	// not care about this return
	return nil
}

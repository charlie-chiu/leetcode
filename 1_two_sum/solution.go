package two_sum

var solution = bruteForce

// first try time complexity: O(n^2)
func bruteForce(nums []int, target int) []int {
	for firstIndex, num := range nums {
		for i2, num2 := range nums[firstIndex+1:] {
			if num+num2 == target {
				return []int{firstIndex, i2 + firstIndex + 1}
			}
		}
	}

	return nil
}

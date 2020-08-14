package max_consecutive_ones

var findMaxConsecutiveOnes func([]int) int = firstTry

func firstTry(nums []int) int {
	const one = 1
	const zero = 0
	var maxConsecutive int
	var windowRight = 0

	for windowLeft := 0; windowRight < len(nums); {
		if nums[windowRight] == one {
			currentLength := windowRight - windowLeft + 1
			if currentLength > maxConsecutive {
				maxConsecutive = currentLength
			}
			windowRight++
		} else if nums[windowRight] == zero {
			windowRight++
			windowLeft = windowRight

		}
	}

	return maxConsecutive
}

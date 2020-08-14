package max_consecutive_ones

var findMaxConsecutiveOnes func([]int) int = secondTry

//sliding window
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

func secondTry(nums []int) int {
	const (
		zero = 0
		one  = 1
	)
	var maxConsecutive, currentMax int

	for _, num := range nums {
		if num == zero {
			currentMax = 0
		} else if num == one {
			currentMax++
		}

		if currentMax > maxConsecutive {
			maxConsecutive = currentMax
		}
	}

	return maxConsecutive
}

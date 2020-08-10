package reverse_integer

import (
	"math"
)

var solution = firstTry

// this solution accepted by leetcode but not working in 32-bit environment
func firstTry(x int) int {
	var output int
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	seq := IntToSlice(x, []int{})

	for digit := 0; digit < len(seq); digit++ {
		output += seq[digit] * int(math.Pow10(digit))
	}

	if output > math.MaxInt32 || output < math.MinInt32 {
		return 0
	}

	return output
}

func IntToSlice(n int, sequence []int) []int {
	if n == 0 {
		return sequence
	}

	i := n % 10
	sequence = append([]int{i}, sequence...)
	return IntToSlice(n/10, sequence)
}

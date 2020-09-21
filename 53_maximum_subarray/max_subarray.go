package maximum_subarray

import (
	"math"
)

// brute force
// time complexity O(n^3)
// got TLE on leetcode
func maxSubArray(nums []int) int {
	var maxSum int = math.MinInt32

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			currSum := sum(nums[i:j])
			if currSum > maxSum {
				maxSum = currSum
			}
		}
	}

	return maxSum
}

func sum(nums []int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}

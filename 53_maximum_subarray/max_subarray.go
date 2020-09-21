package maximum_subarray

import "math"

/*
	time complexity: O(n)
	is this kadane's algorithm?
*/
func maxSubArray(nums []int) int {
	var maxSum int = math.MinInt32
	var currSum int

	for _, num := range nums {
		if currSum < 0 {
			currSum = num
		} else {
			currSum += num
		}

		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return maxSum
}

// improved brute force
// 84ms (5.96%), 3.3MB on leetcode
//func maxSubArray(nums []int) int {
//	var maxSum int = math.MinInt32
//	var currSum int
//
//	for i := 0; i < len(nums); i++ {
//		currSum = 0
//		for j := i; j < len(nums); j++ {
//			currSum += nums[j]
//			if currSum > maxSum {
//				maxSum = currSum
//			}
//		}
//	}
//
//	return maxSum
//}

// brute force
// time complexity O(n^3)
// got TLE on leetcode
//func maxSubArray(nums []int) int {
//	var maxSum int = math.MinInt32
//
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j <= len(nums); j++ {
//			currSum := sum(nums[i:j])
//			if currSum > maxSum {
//				maxSum = currSum
//			}
//		}
//	}
//
//	return maxSum
//}
//
//func sum(nums []int) (sum int) {
//	for _, num := range nums {
//		sum += num
//	}
//	return
//}

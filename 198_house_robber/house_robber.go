package house_robber

//iterative with 2 variables
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var prev1, prev2 int
	for _, num := range nums {

		oldPrev1 := prev1
		prev1 = max(prev1, num+prev2)
		prev2 = oldPrev1
	}

	return prev1
}

// recursive
//func rob(nums []int) int {
//	var memo = make(map[int]int)
//	return helper(nums, len(nums)-1, memo)
//}
//
//func helper(nums []int, n int, memo map[int]int) int {
//	if n < 0 {
//		return 0
//	}
//	if loot, ok := memo[n]; ok {
//		return loot
//	}
//
//	prevHouse := helper(nums, n-1, memo)
//	memo[n-1] = prevHouse
//	prevPrevHouse := helper(nums, n-2, memo)
//	memo[n-2] = prevPrevHouse
//	currHouseLoot := nums[n]
//	return max(prevHouse, prevPrevHouse+currHouseLoot)
//}
//
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

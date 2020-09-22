package climbing_stairs

// leetcode approach1 : brute force
// bottom up
func climbStairs(n int) int {
	return helper(0, n)
}

func helper(i, n int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}

	return helper(i+1, n) + helper(i+2, n)
}

// dynamic programming with memoization
// 0ms, 2MB
//var memoization = map[int]int{
//	1: 1,
//	2: 2,
//}
//
//func climbStairs(n int) int {
//	if ways, ok := memoization[n]; ok {
//		return ways
//	}
//
//	ways := climbStairs(n-1) + climbStairs(n-2)
//	memoization[n] = ways
//
//	return ways
//}

//got TLE on leetcode
//func climbStairs(n int) int {
//	if n == 1 {
//		return 1
//	}
//
//	if n == 2 {
//		return 2
//	}
//
//	return climbStairs(n-1) + climbStairs(n-2)
//}

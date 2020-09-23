package unique_paths

// dp solution?
func uniquePaths(m int, n int) int {
	var grid [101][101]int

	for i := 1; i <= n; i++ {
		grid[1][i] = 1
	}
	for i := 1; i <= m; i++ {
		grid[i][1] = 1
	}

	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	return grid[m][n]
}

// first try, improved with memoization
// 0ms on leetcode
//type Key struct {
//	m, n int
//}
//
//func uniquePaths(m int, n int) int {
//	var memoization = make(map[Key]int)
//	//base case
//	memoization[Key{1, 1}] = 1
//	return helper(Key{m, n}, memoization)
//}
//
//func helper(k Key, m map[Key]int) int {
//	if ways, ok := m[k]; ok {
//		return ways
//	}
//	// on edge, only one way
//	if k.m == 1 || k.n == 1 {
//		return 1
//	}
//
//	fromAbove := helper(Key{k.m - 1, k.n}, m)
//	m[Key{k.m - 1, k.n}] = fromAbove
//	fromLeft := helper(Key{k.m, k.n - 1}, m)
//	m[Key{k.m, k.n - 1}] = fromLeft
//	return fromAbove + fromLeft
//}

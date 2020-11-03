package tiling_problem

//https://www.geeksforgeeks.org/tiling-problem/
// count tile ways
// in 2 x n board

func count(n int) (ways int) {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return count(n-1) + count(n-2)
}

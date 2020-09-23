package unique_paths_ii

// 0ms
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	finishM := len(obstacleGrid)
	finishN := len(obstacleGrid[0])
	var path [101][101]int
	// dummy start, path[1][0] work, too
	path[0][1] = 1

	for m := 1; m <= finishM; m++ {
		for n := 1; n <= finishN; n++ {
			if obstacleGrid[m-1][n-1] == 1 {
				path[m][n] = 0
			} else {
				path[m][n] = path[m-1][n] + path[m][n-1]
			}
		}
	}

	return path[finishM][finishN]
}

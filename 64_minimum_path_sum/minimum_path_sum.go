package minimum_path_sum

func minPathSum(grid [][]int) int {
	numberOfRows := len(grid)
	numberOfColumns := len(grid[0])

	minPathSum := make([][]int, numberOfRows+1)
	for i := range minPathSum {
		minPathSum[i] = make([]int, numberOfColumns+1)
	}

	for m := 1; m <= numberOfRows; m++ {
		for n := 1; n <= numberOfColumns; n++ {
			currPath := grid[m-1][n-1]
			if m == 1 && n == 1 {
				// start point
				minPathSum[m][n] = currPath
			} else if m == 1 {
				// first row
				minPathSum[m][n] = currPath + minPathSum[m][n-1]
			} else if n == 1 {
				// first column
				minPathSum[m][n] = currPath + minPathSum[m-1][n]
			} else {
				// other cells
				minPathSum[m][n] = currPath + min(minPathSum[m][n-1], minPathSum[m-1][n])
			}
		}
	}

	return minPathSum[numberOfRows][numberOfColumns]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

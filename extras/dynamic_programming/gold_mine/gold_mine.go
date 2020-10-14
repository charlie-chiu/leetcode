package gold_mine

//https://www.geeksforgeeks.org/gold-mine-problem/

func getMaxGold(mine [][]int) (profit int) {
	var numberOfRow = len(mine)
	var numberOfCol = len(mine[0])
	var maxAmount = make([][]int, numberOfRow)
	for i := range maxAmount {
		maxAmount[i] = make([]int, numberOfCol)
	}

	for column := 0; column < numberOfCol; column++ {
		for row := 0; row < numberOfRow; row++ {
			// first column
			if column == 0 {
				maxAmount[row][column] = mine[row][0]
				continue
			}

			if row == 0 {
				// first row
				maxAmount[row][column] = mine[row][column] + max(maxAmount[row][column-1], maxAmount[row+1][column-1])
			} else if row == numberOfRow-1 {
				// last row
				maxAmount[row][column] = mine[row][column] + max(maxAmount[row][column-1], maxAmount[row-1][column-1])
			} else {
				maxAmount[row][column] = mine[row][column] + max(maxAmount[row][column-1], maxAmount[row+1][column-1], maxAmount[row-1][column-1])
			}

			if maxAmount[row][column] > profit {
				profit = maxAmount[row][column]
			}
		}
	}

	return
}

func max(args ...int) (max int) {
	for _, arg := range args {
		if arg > max {
			max = arg
		}
	}
	return
}

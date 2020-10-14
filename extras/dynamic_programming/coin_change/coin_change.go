package coin_change

//https://www.geeksforgeeks.org/coin-change-dp-7/?ref=lbp

// with table
//func countWays(values []int, targetSum int) (ways int) {
//	// initial table
//	var table = make([][]int, targetSum+1)
//	for currSum := range table {
//		table[currSum] = make([]int, len(values))
//	}
//
//	// we have 1 way to complete sum 0 (include 0 coins)
//	for i := range table[0] {
//		table[0][i] = 1
//	}
//
//	for currTarget := 1; currTarget < targetSum+1; currTarget++ {
//		var countInclude, countExclude int
//		for valueIndex := 0; valueIndex < len(values); valueIndex++ {
//			// count of solutions including curr coin value
//			currValue := values[valueIndex]
//			if currTarget-currValue >= 0 {
//				countInclude = table[currTarget-currValue][valueIndex]
//			} else {
//				// no solutions while curr value greater than sum
//				countInclude = 0
//			}
//
//			// count of solutions excluding curr coin value
//			// index == 0 means no any coin after exclude curr coin
//			if valueIndex > 0 {
//				countExclude = table[currTarget][valueIndex-1]
//			} else {
//				countExclude = 0
//			}
//
//			table[currTarget][valueIndex] = countInclude + countExclude
//		}
//	}
//
//	return table[targetSum][len(values)-1]
//}

//basic recursive without memoization
func countWays(coinValues []int, sum int) (ways int) {
	// base cases
	if sum == 0 {
		return 1
	}

	if sum < 0 {
		return 0
	}

	// no solutions while no coins
	if len(coinValues) == 0 {
		return 0
	}

	lastCoinValue := coinValues[len(coinValues)-1]
	countIncludingLastCoin := countWays(coinValues, sum-lastCoinValue)
	countExcludingLastCoin := countWays(coinValues[:len(coinValues)-1], sum)

	return countIncludingLastCoin + countExcludingLastCoin
}

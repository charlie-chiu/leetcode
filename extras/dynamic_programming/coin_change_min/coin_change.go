package coin_change_min

import (
	"math"
	"sort"
)

// https://www.geeksforgeeks.org/find-minimum-number-of-coins-that-make-a-change/?ref=lbp
// same as leetcode 322

//dp - tabulation - bottom-up
func coinChange(coins []int, target int) (minCoins int) {
	sort.Ints(coins)
	//if target == 0 {
	//	return 0
	//}

	// initial table
	var table = make([][]int, len(coins))
	for coinValue := range table {
		table[coinValue] = make([]int, target+1)
	}

	for amount := 0; amount <= target; amount++ {
		for coinIndex := 0; coinIndex < len(coins); coinIndex++ {
			if amount == 0 {
				// zero coin made up amount 0
				table[coinIndex][amount] = 0
				continue
			}

			coinValue := coins[coinIndex]

			if coinValue > amount {
				if coinIndex == 0 {
					table[coinIndex][amount] = -1
				} else {
					table[coinIndex][amount] = table[coinIndex-1][amount]
				}
			}

			if coinValue == amount {
				table[coinIndex][amount] = 1
			}

			if coinValue < amount {
				table[coinIndex][amount] = table[coinIndex][amount-coinValue] + 1
			}
		}
	}

	minCoins = math.MaxInt64
	for i := 0; i < len(coins); i++ {
		//log.Println(table[i])
		if table[i][target] >= -1 && table[i][target] < minCoins {
			minCoins = table[i][target]
		}
	}
	return
}

// basic recursive
//func coinChange(coins []int, amount int) (minCoins int) {
//	// base case
//	if amount == 0 {
//		return 0
//	}
//
//	minCoins = math.MaxInt64
//	for _, value := range coins {
//		if value <= amount {
//			tempCoins := coinChange(coins, amount -value) + 1
//			if tempCoins < minCoins {
//				minCoins = tempCoins
//			}
//		}
//	}
//
//	return
//}

// greedy algorithm, won't work at some cases
//func coinChange(coins []int, amount int) (minWays int) {
//	sort.Ints(coins)
//	coinIndex := len(coins) - 1
//	for amount > 0 {
//		var biggestValue = coins[coinIndex]
//
//		if amount < biggestValue {
//			coinIndex--
//			continue
//		}
//
//		for amount >= biggestValue {
//			amount -= biggestValue
//			minWays++
//		}
//	}
//
//	return
//}

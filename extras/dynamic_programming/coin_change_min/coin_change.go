package coin_change_min

import (
	"log"
	"sort"
)

// https://www.geeksforgeeks.org/find-minimum-number-of-coins-that-make-a-change/?ref=lbp
// same as leetcode 322

//dp - tabulation - bottom-up
func coinChange(coins []int, amount int) (minCoins int) {
	if amount == 0 {
		return 0
	}
	sort.Ints(coins)

	var table = make([][]int, amount+1)
	for i := range table {
		table[i] = make([]int, len(coins))
	}

	for target := 0; target <= amount; target++ {
		for denoIndex := 0; denoIndex < len(coins); denoIndex++ {
			if target == 0 {
				// 0 coin made up target amount 0
				table[target][denoIndex] = 0
				continue
			}

			denomination := coins[denoIndex]

			//log.Println(table[target])
			if denomination > target {
				if denoIndex == 0 {
					// no smaller denomination
					table[target][denoIndex] = -1
				} else {
					table[target][denoIndex] = table[target][denoIndex-1]
				}
			}

			if denomination == target {
				table[target][denoIndex] = 1
			}

			if denomination < target {
				table[target][denoIndex] = table[target-denomination][denoIndex] + 1
			}
		}
	}

	for i, t := range table {
		log.Println(i, t)
	}

	sort.Ints(table[amount])
	if table[amount][0] > 0 {
		return table[amount][0]
	} else {
		return -1
	}
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

package best_time_to_buy_and_sell_stock

import "math"

// one pass from leetcode solution
func maxProfit(prices []int) int {
	var minPrice = math.MaxInt64
	var maxProfit int
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}

// brute force
//
//func maxProfit(prices []int) int {
//	var maxProfit int
//	for i, buyingPrice := range prices {
//		for j, sellingPrice := range prices {
//			if j <= i {
//				continue
//			}
//
//			profit := sellingPrice - buyingPrice
//			if profit > 0 && profit > maxProfit {
//				maxProfit = profit
//			}
//		}
//	}
//
//	return maxProfit
//}

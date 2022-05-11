package main

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxProfit := 0

	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		} else if v-minPrice > maxProfit {
			maxProfit = v - minPrice
		}
	}

	return maxProfit
}

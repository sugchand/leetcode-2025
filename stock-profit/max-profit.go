package main

import "fmt"

// You are given an array prices where prices[i] is the price of a given stock on the ith day.

// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
func maxProfit(prices []int) int {

	profit := 0
	sellPrice := prices[len(prices)-1]

	lowIdx := 0
	val := prices[0]
	for i, k := range prices {
		if k < val {
			val = k
			lowIdx = i
		}
	}
	return profit
}

func main() {
	fmt.Print(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

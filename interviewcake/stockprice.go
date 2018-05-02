package main

import "fmt"

/*
int[] stockPricesYesterday = new int[] {10, 7, 5, 8, 11, 9};

getMaxProfit(stockPricesYesterday);
// returns 6 (buying for $5 and selling for $11)
*/

func getBruteForceMaxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > max {
				max = profit
			}
		}
	}
	return max
}

func getMaxProfit(prices []int) int {
	//TODO: throw exception in case arr length is less than 2
	minPrice := prices[0]
	max := prices[1] - prices[0]
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - minPrice
		if profit > max {
			max = profit
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return max
}

func main() {
	stockPrices := []int{10, 7, 5, 8, 11, 9}
	fmt.Println(fmt.Sprintf("max stock profit: %d", getMaxProfit(stockPrices)))
}
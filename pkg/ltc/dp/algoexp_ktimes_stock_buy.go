package dp

import (
// "fmt"
)


// O(nk) time, O(n) memory
// optimized the inner for loop to find max from prev. days.
// optimized to store only profit from (t-1) buy/sell.
func MaxProfitWithKTransactions3(prices []int, k int) int {
	/*
		// profit[t][d] -> profit at day 'd' bought at max 't' times
		// profit[t][d] = max {
								profit[t-1][d] // we don't buy/sell again --> (1)
								prices[d] + max {
									for x = 0 to d-1
										 -prices[x] + profit[t-1][x]
								}
								prices[d] - prices[d-1] + profit[t-1][d-1] // buy on d-1 + prev. profit
								prices[d] - prices[d-2] + profit[t-1][d-2]
								....
								prices[d] = prices[0] + profit[t-1][0]


	*/
	n := len(prices)
	if n == 0 {
		return 0
	}

	prevProfit := make([]int, n)
	currProfit := make([]int, n)

	for t := 1; t <= k; t++ {
		maxNow := -100000
		for d := 1; d < n; d++ {
			// sell on day 'd', buy on 'x', maxNow is aggregating all previous x's.
			maxNow = max(maxNow, -prices[d-1]+prevProfit[d-1])
			// don't buy/sell today, carry over profit from yesterday.
			currProfit[d] = max(currProfit[d-1], maxNow + prices[d])
		}
		tmp := prevProfit
		prevProfit = currProfit
		currProfit = tmp
	}

	return prevProfit[n-1]
}

// O(nk) time, O(nk) memory
// optimized the inner for loop to find max from prev. days.
func MaxProfitWithKTransactions2(prices []int, k int) int {
	/*
		// profit[t][d] -> profit at day 'd' bought at max 't' times
		// profit[t][d] = max {
								profit[t-1][d] // we don't buy/sell again --> (1)
								prices[d] + max {
									for x = 0 to d-1
										 -prices[x] + profit[t-1][x]
								}
								prices[d] - prices[d-1] + profit[t-1][d-1] // buy on d-1 + prev. profit
								prices[d] - prices[d-2] + profit[t-1][d-2]
								....
								prices[d] = prices[0] + profit[t-1][0]


	*/
	n := len(prices)
	if n == 0 {
		return 0
	}

	// k+1, 0th row is used as base case
	profits := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		profits[i] = make([]int, n)
	}

	for t := 1; t <= k; t++ {
		maxNow := -100000
		for d := 1; d < n; d++ {
			// sell on day 'd', buy on 'x', maxNow is aggregating all previous x's.
			maxNow = max(maxNow, -prices[d-1]+profits[t-1][d-1])
			// don't buy/sell today, carry over profit from yesterday.
			profits[t][d] = max(profits[t][d-1], maxNow + prices[d])
		}
		
	}

	return profits[k][n-1]
}

// O(n^2k) time, O(nk) memory
func MaxProfitWithKTransactions1(prices []int, k int) int {
	/*
	   // profit[t][d] -> profit at day 'd' bought at max 't' times
	   // profit[t][d] = max {
	                           profit[t-1][d] // we don't buy/sell again --> (1)
	                           prices[d] + max {
	                               for x = 0 to d-1
	                                    -prices[x] + profit[t-1][x]
	                           }
	                           prices[d] - prices[d-1] + profit[t-1][d-1] // buy on d-1 + prev. profit
	                           prices[d] - prices[d-2] + profit[t-1][d-2]
	                           ....
	                           prices[d] = prices[0] + profit[t-1][0]


	*/
	n := len(prices)
	if n == 0 {
		return 0
	}

	// k+1, 0th row is used as base case
	profits := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		profits[i] = make([]int, n)
	}

	for t := 1; t <= k; t++ {
		for d := 1; d < n; d++ {
			// don't buy/sell today, carry over profit from yesterday.
			maxNow := profits[t][d-1]
			// sell on day 'd', buy on 'x'
			for x := 0; x < d; x++ {
				profitOnD := prices[d] - prices[x] + profits[t-1][x]
				maxNow = max(maxNow, profitOnD)
			}
			profits[t][d] = maxNow
		}
		// fmt.Printf("profits[%d] is %v\n", t, profits[t])
	}

	return profits[k][n-1]
}

// func max(a, b int) int {
//     if a > b {return a}
//     return b
// }

package dp

func coinChange(coins []int, amount int) int {
	// val := coinrec(coins, amount)
	// val := cointd(coins, amount)
	val := coinchiter(coins, amount)
	if val == ccinf {
		return -1
	}
	return val
}

func coinchiter(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = ccinf
	}
	dp[0] = 0

	for am := 1; am < amount+1; am++ {
		for i := 0; i < len(coins); i++ {
			if am-coins[i] >= 0 {
				dp[am] = min322(dp[am], 1+dp[am-coins[i]])
			}
		}
	}

	return dp[amount]
}

var ccinf = 100000 // 10^5

func cointd(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = -1
	}
	return coinmem1(coins, dp, amount)
}

func coinmem1(coins, dp []int, amount int) int {
	if amount < 0 {
		return ccinf
	}
	if dp[amount] != -1 {
		return dp[amount]
	}
	if amount == 0 {
		return 0
	}
	ans := ccinf
	for i := 0; i < len(coins); i++ {
		ans = min322(ans, 1+coinmem1(coins, dp, amount-coins[i]))
	}
	dp[amount] = ans
	return dp[amount]
}

func coinrec(coins []int, amount int) int {
	if amount < 0 {
		return ccinf
	}

	if amount == 0 {
		return 0
	}

	ans := ccinf
	for i := 0; i < len(coins); i++ {
		ans = min322(ans, 1+coinrec(coins, amount-coins[i]))
	}

	return ans
}

func min322(a, b int) int {
	if a < b {
		return a
	}
	return b
}

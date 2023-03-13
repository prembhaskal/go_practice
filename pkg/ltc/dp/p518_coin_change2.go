package dp

func change(amount int, coins []int) int {
	// fmt.Println("\n\n***************************\n\n")
	// return f1mem(amount, coins)
	// return frec(amount, coins)
	// return fiter(amount, coins)
	// return fiterlessmem(amount, coins)
	return fiteropt(amount, coins)
}

// optimized version.
func fiteropt(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1 // base case

	for i := 1; i < len(coins)+1; i++ {
		for am := 1; am < amount+1; am++ {
			// val := dp[am]
			if am-coins[i-1] >= 0 {
				dp[am] = dp[am] + dp[am-coins[i-1]]
			}
			// dp[am] = val
		}
	}
	return dp[amount]
}

func fiterlessmem(amount int, coins []int) int {
	// dp[i][j] = dp[i][j - coins[i-1]] + dp[i-1][j]
	// dp depends on self previous index or previous dp only.
	prev := make([]int, amount+1)
	prev[0] = 1

	for i := 1; i < len(coins)+1; i++ {
		curr := make([]int, amount+1)
		curr[0] = 1
		for am := 1; am < amount+1; am++ {
			if am-coins[i-1] >= 0 {
				curr[am] = curr[am-coins[i-1]] + prev[am]
			} else {
				curr[am] = prev[am]
			}
		}
		prev = curr
	}
	return prev[amount]
}

func fiter(amount int, coins []int) int {
	// dp[i][j] = count of combinations with 1st 'i' coins to make 'j' amount.
	dp := make([][]int, 0)
	for i := 0; i < len(coins)+1; i++ {
		row := make([]int, amount+1)
		dp = append(dp, row)
	}

	// base case
	// with no coins, no combinations. dp[0][x] is set to ZERO automatically.
	// with some coin > 0, if we reach amount = 0, then 1 combination is possible.

	for i := 1; i < len(coins)+1; i++ {
		dp[i][0] = 1
	}

	//   coins[i-1] because index is 0 based, the terminology is 1 based.
	// dp[i][j] = dp[i][j - coins[i-1]] + dp[i-1][j]
	for i := 1; i < len(coins)+1; i++ {
		for am := 1; am < amount+1; am++ {
			if am-coins[i-1] >= 0 {
				dp[i][am] = dp[i][am-coins[i-1]] + dp[i-1][am]
			} else {
				dp[i][am] = dp[i-1][am]
			}
		}
	}

	return dp[len(coins)][amount]
}

func frec(amount int, coins []int) int {
	// return f2(len(coins), amount, coins)
	mem := make([][]int, 0)
	for n := 0; n < len(coins)+1; n++ {
		mem = append(mem, make([]int, amount+1))
	}
	for i := 0; i < len(coins)+1; i++ {
		for j := 0; j < amount+1; j++ {
			mem[i][j] = -1
		}
	}
	return f2mem(len(coins), amount, coins, mem)
}

func f2mem(n, amount int, coins []int, mem [][]int) int {
	if amount < 0 {
		return 0
	}
	if amount == 0 {
		return 1
	}
	if n == 0 {
		return 0
	}
	if mem[n][amount] != -1 {
		return mem[n][amount]
	}
	// classic choose/don't choose pattern in work here.
	// choose (n-1)th coin, don't choose (n-1)th coin.
	// f(n, amount - coins[n-1]...) means is , we use one coin with value 'coins[n-1]' and check combination with remaining amount
	// f(n-1, amount, ...) means is , we don't use this coin and see if how many combinations we can make with this amount
	mem[n][amount] = f2mem(n, amount-coins[n-1], coins, mem) +
		f2mem(n-1, amount, coins, mem)
	return mem[n][amount]
}

// ways to make amount using 1st n coins
func f2(n, amount int, coins []int) int {
	if amount < 0 {
		return 0
	}
	if amount == 0 {
		return 1
	}

	if n == 0 { // if no coins, then no combination.
		return 0
	}

	// n-1 because of 0 based index, but this is actually the nth coin
	// make amount-coins[n-1] with 1st n coins + make amount with 1st n-1 coins
	return f2(n, amount-coins[n-1], coins) + f2(n-1, amount, coins)
}

// var ccinf = 10000 // 10^4

func f1mem(amount int, coins []int) int {
	mem := make([][]int, 0)
	for i := 0; i < amount+1; i++ {
		row := make([]int, len(coins))
		mem = append(mem, row)
	}
	for i := 0; i < amount+1; i++ {
		for j := 0; j < len(coins); j++ {
			mem[i][j] = -1
		}
	}

	return f1(amount, coins, 0, mem)
}

// this works, but it is overly complicated.
func f1(amount int, coins []int, coinidx int, mem [][]int) int {
	// fmt.Printf("enter amount: %d, coinidx: %d\n", amount, coinidx)
	if amount < 0 {
		return 0
	}
	if mem[amount][coinidx] != -1 {
		return mem[amount][coinidx]
	}

	if amount == 0 {
		return 1
	}

	total := 0
	for i := coinidx; i < len(coins); i++ {
		total = total + f1(amount-coins[i], coins, i, mem)
	}
	// fmt.Printf("exit amount: %d, combinations: %d\n", amount, total)
	mem[amount][coinidx] = total
	return total
}

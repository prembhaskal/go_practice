package dp

func change(amount int, coins []int) int {
	// fmt.Println("\n\n***************************\n\n")
	// return f1mem(amount, coins)
	return frec(amount, coins)
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

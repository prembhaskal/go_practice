package dp

func minCostClimbingStairs(cost []int) int {
	// return minstaircost(cost)
	return minstaircost2(cost)
}

// bottom up
func minstaircost(cost []int) int {
	// dp[i] = min(cost[i] + dp[i-1], cost[i] + dp[i-2])
	cost = append(cost, 0)
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(cost[i]+dp[i-1], cost[i]+dp[i-2])
	}
	return dp[len(cost)-1]
}

// top down
func minstaircost2(cost []int) int {
	// f(i)  = min(i + f(i-1), i + f(i-2))
	cost = append(cost, 0)
	mem := make([]int, len(cost))
	for k := range mem {
		mem[k] = -1
	}
	return minstaircostrec(len(cost)-1, cost, mem)
}

func minstaircostrec(i int, cost []int, mem []int) int {
	if i < 0 {
		return 0
	}
	if mem[i] != -1 {
		return mem[i]
	}
	prev := cost[i] + minstaircostrec(i-1, cost, mem)
	pprev := cost[i] + minstaircostrec(i-2, cost, mem)
	mem[i] = min(prev, pprev)
	// fmt.Printf("minval at i:%d is %d\n", i, minval)
	return mem[i]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

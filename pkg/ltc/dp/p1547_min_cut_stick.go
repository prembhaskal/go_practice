package dp

func minCost(n int, cuts []int) int {
	dp := make(map[pair1547]int)
	return minCostRec(0, n, cuts, dp)
}

// TODO - optimize by using indices of cuts instead of string to run it faster
// DP[i][j] = min cost for cutting between cuts[i] and cuts[j] (both not included ??)
// which is equal to (cuts[j] - cuts[i]) + DP[i][k] + DP[k][j] where k ranges from i + 1 to j - 1
func minCostRec(start, end int, cuts []int, dp map[pair1547]int) int {
	if start >= end {
		return 0
	}

	if val, ok := dp[pair1547{start, end}]; ok {
		return val
	}

	cutfound := false
	mincost := 10000000
	for _, cut := range cuts {
		if cut > start && cut < end {
			cutfound = true
			mincost = min1547(mincost, minCostRec(start, cut, cuts, dp)+minCostRec(cut, end, cuts, dp)+(end-start))
		}
	}

	if !cutfound {
		mincost = 0
	}
	dp[pair1547{start, end}] = mincost
	return dp[pair1547{start, end}]
}

type pair1547 struct {
	st int
	en int
}

func min1547(a, b int) int {
	if a < b {
		return a
	}
	return b
}

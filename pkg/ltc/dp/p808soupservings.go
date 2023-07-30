package dp

func soupServings(n int) float64 {
	return soupServingsBotUp(n)
}

// Approach - DP bottom up, from each stage, we have 4 next stages. we update next stage
// next stage = next stage + (current_stage) * 0.25
// A empty first is tracked in separately
// for n  5000, it always give 1.0 as found by testing.
func soupServingsBotUp(n int) float64 {
	dp := make(map[pair808]float64)
	aemptyFirst = make(map[int]float64)

	dp[pair808{n, n}] = 1.0

	if n > 5000 {
		return 1.0
	}

	for ac := n; ac > 0; {
		for bc := n; bc > 0; {
			updateDP(ac-100, bc, ac, bc, dp)
			updateDP(ac-75, bc-25, ac, bc, dp)
			updateDP(ac-50, bc-50, ac, bc, dp)
			updateDP(ac-25, bc-75, ac, bc, dp)
			bc = bc - 25
		}

		ac = ac - 25
	}

	// total a emptyFirst
	var aempty float64
	for _, v := range aemptyFirst {
		aempty += v
	}

	// fmt.Printf("aemptyFirst is %f\n", aempty)
	// fmt.Printf("both empty is %f\n", dp[pair{0,0}])

	return aempty + 0.5*dp[pair808{0, 0}]
}

var aemptyFirst map[int]float64

func updateDP(na, nb, a, b int, dp map[pair808]float64) {
	na = max808(na, 0)
	nb = max808(nb, 0)
	dp[pair808{na, nb}] += (dp[pair808{a, b}] * 0.25)
	if na == 0 && nb > 0 {
		aemptyFirst[nb] = dp[pair808{na, nb}]
	}

	// fmt.Printf("na: %d, nb: %d, a: %d, b :%d, prob: %f\n", na, nb, a, b, dp[pair{na, nb}])

}

func max808(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type pair808 struct {
	a int
	b int
}

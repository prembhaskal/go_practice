package recurse

func climbStairs(n int) int {
	// ways[n]  = ways[n-1] + ways[n-2]

	// return climbStairsDP(n)
	// return climbStairsIter(n)

	return climbStairsRec(n, make(map[int]int))
}

func climbStairsRec(n int, mem map[int]int) int {
	if n == 0 || n == 1 {
		return 1
	}

	if val, ok := mem[n]; ok {
		return val
	}
	mem[n] = climbStairsRec(n-1, mem) + climbStairsRec(n-2, mem)
	return mem[n]
}

func climbStairsDP(n int) int {
	ways := make([]int, n+1)

	ways[0] = 1 // one way to not climb stairs yet
	ways[1] = 1 // one way to reach 1st stair

	for i := 2; i < n+1; i++ {
		ways[i] = ways[i-1] + ways[i-2]
	}
	return ways[n]
}

func climbStairsIter(n int) int {
	a2 := 1
	a1 := 1

	if n == 1 {
		return 1
	}

	for i := 2; i < n+1; i++ {
		c := a1 + a2

		// next round
		a2 = a1
		a1 = c
	}
	return a1
}

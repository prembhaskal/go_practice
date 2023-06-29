package dp

func longestCommonSubsequence(text1 string, text2 string) int {
	// return lcsusingrec(text1, text2)
	// return lcsusingbotup(text1, text2)
	// return lcsusingbotupspace(text1, text2)
	return lcsusingbotupspacelessCond(text1, text2)
}

func lcsusingbotupspacelessCond(s1, s2 string) int {
	n := len(s2)
	// additional row to remove extra if else checks
	prev := make([]int, n+1)
	curr := make([]int, n+1)

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				curr[j] = 1 + prev[j-1]
			} else {
				curr[j] = max1143(prev[j], curr[j-1])
			}
		}
		// trying to make the thing run faster by not creating new array.
		prev, curr = curr, prev
		for x := range prev { // use range loop gives bit faster than loop
			curr[x] = 0
		}
	}

	return prev[n]
}

func lcsusingbotupspace(s1, s2 string) int {
	m := len(s1)
	n := len(s2)
	prev := make([]int, n)
	curr := make([]int, n)

	A := []rune(s1)
	B := []rune(s2)

	// base case
	if A[0] == B[0] {
		prev[0] = 1
	}

	for j := 1; j < n; j++ {
		if B[j] == A[0] || prev[j-1] == 1 {
			prev[j] = 1
		}
	}

	// fmt.Printf("prev: %v\n", prev)

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if A[i] == B[j] {
				if j == 0 {
					curr[j] = 1
				} else {
					curr[j] = 1 + prev[j-1]
				}
			} else {
				if j == 0 {
					curr[j] = prev[j]
				} else {
					curr[j] = max1143(prev[j], curr[j-1])
				}
			}
		}
		// fmt.Printf("i:%d\n", i)
		// fmt.Printf("prev: %v\n", prev)
		// fmt.Printf("curr: %v\n", curr)
		// prev = curr
		// curr = make([]int, n)
		// avoid mem allocation by resetting values.
		for x := 0; x < len(prev); x++ {
			prev[x] = curr[x]
			curr[x] = 0
		}
	}

	return prev[n-1]
}

func lcsusingbotup(s1, s2 string) int {
	m := len(s1)
	n := len(s2)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	A := []rune(s1)
	B := []rune(s2)

	// base case
	if A[0] == B[0] {
		dp[0][0] = 1
	}

	for i := 1; i < m; i++ {
		if A[i] == B[0] || dp[i-1][0] == 1 {
			dp[i][0] = 1
		}
	}

	for j := 1; j < n; j++ {
		if A[0] == B[j] || dp[0][j-1] == 1 {
			dp[0][j] = 1
		}
	}

	// fmt.Printf("dp:%v\n", dp)

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if A[i] == B[j] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max1143(dp[i-1][j], dp[i][j-1])
			}
			// fmt.Printf("i: %d, j: %d, dp: %v\n", i, j, dp)
		}
	}

	return dp[m-1][n-1]
}

func lcsusingrec(s1, s2 string) int {
	m := len(s1)
	n := len(s2)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	return lcsrec([]rune(s1), []rune(s2), len(s1)-1, len(s2)-1, dp)
}

func lcsrec(s1, s2 []rune, i, j int, dp [][]int) int {
	if i < 0 || j < 0 {
		return 0
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}

	if s1[i] == s2[j] {
		dp[i][j] = 1 + lcsrec(s1, s2, i-1, j-1, dp)
		return dp[i][j]
	}

	dp[i][j] = max1143(lcsrec(s1, s2, i, j-1, dp), lcsrec(s1, s2, i-1, j, dp))
	return dp[i][j]
}

func max1143(a, b int) int {
	if a > b {
		return a
	}
	return b
}

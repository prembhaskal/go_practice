package dp

func minDistance(word1 string, word2 string) int {
	return mindistdp([]rune(word1), []rune(word2))
}

func mindistdp(w1, w2 []rune) int {
	m := len(w1)
	n := len(w2)
	dp := make([][]int, 0)

	// 1 extra to account for base cases.
	for i := 0; i < m+1; i++ {
		row := make([]int, n+1)
		dp = append(dp, row)
	}

	// 1st column
	for j := 0; j < n+1; j++ {
		dp[0][j] = j
	}
	// 1st row
	for i := 1; i < m+1; i++ {
		dp[i][0] = i
	}

	// rest
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			dp[i][j] = min72(min72(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+diff72(w1[i-1], w2[j-1])) // i-1,j-1 because actual string are still 0-index
		}
	}

	return dp[m][n]
}

func diff72(a, b rune) int {
	if a == b {
		return 0
	}
	return 1
}

func min72(a, b int) int {
	if a < b {
		return a
	}
	return b
}

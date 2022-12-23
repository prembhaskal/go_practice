package dp

func longestCommonSubsequence(text1 string, text2 string) int {
	return lcsusingrec(text1, text2)
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

package dp

func shortestCommonSupersequence(str1 string, str2 string) string {
	// find LCS first
	return scsrec(str1, str2)
}

func scsrec(s1, s2 string) string {
	dp := make([][]int, 0)
	for i := 0; i < len(s1)+1; i++ {
		row := make([]int, len(s2)+1)
		dp = append(dp, row)
	}
	s1r := []rune(s1)
	s2r := []rune(s2)
	fscsrec(s1r, s2r, len(s1), len(s2), dp)

	// fmt.Println("\n************\n")
	// for i := 0; i < len(s1r)+1; i++ {
	//     for j := 0; j < len(s2r)+1; j++ {
	//         fmt.Printf(" %d ", dp[i][j])
	//     }
	//     fmt.Println()
	// }
	// fmt.Println("\n************\n")

	// recreate string
	scs := make([]rune, 0)
	i := len(s1r)
	j := len(s2r)
	for i > 0 && j > 0 {
		if s1r[i-1] == s2r[j-1] {
			scs = append(scs, s1r[i-1]) // or s2r[j] , both will be same
			i--
			j--
		} else if dp[i][j] == 1+dp[i-1][j] {
			scs = append(scs, s1r[i-1])
			i--
		} else {
			// last option
			scs = append(scs, s2r[j-1])
			j--
		}
	}
	for ; i > 0; i-- {
		scs = append(scs, s1r[i-1])
	}
	for ; j > 0; j-- {
		scs = append(scs, s2r[j-1])
	}

	return string(rev(scs))
}

func rev(s []rune) []rune {
	i := 0
	j := len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return s
}

// f(i,j) means scs with 'i' chars from S1 and 'j' chars from S2
func fscsrec(s1, s2 []rune, i, j int, dp [][]int) int {
	// fmt.Printf("i: %d, j: %d\n", i, j)
	// base cases, if no chars in S1, then SCS(0,j) = j
	if i == 0 {
		dp[i][j] = j
		return dp[i][j]
	}
	// SCS(i,0) = i
	if j == 0 {
		dp[i][j] = i
		return dp[i][j]
	}
	if dp[i][j] != 0 {
		return dp[i][j]
	}

	if s1[i-1] == s2[j-1] {
		dp[i][j] = 1 + fscsrec(s1, s2, i-1, j-1, dp)
	} else {
		dp[i][j] = 1 + min1092(fscsrec(s1, s2, i, j-1, dp), fscsrec(s1, s2, i-1, j, dp))
	}
	// fmt.Printf("return val i: %d, j: %d, val: %d\n", i, j, dp[i][j])
	return dp[i][j]
}

func min1092(a, b int) int {
	if a < b {
		return a
	}
	return b
}

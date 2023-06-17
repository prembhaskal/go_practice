package str

func countSubstrings(s string) int {
	// return countSubstringswithdp(s)
	return countFromcentre(s)
}

func countFromcentre(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		total += countPalindrome(i, i, len(s), s)
		total += countPalindrome(i, i+1, len(s), s)
	}
	return total
}

func countPalindrome(start, end, n int, s string) int {
	count := 0
	if start == end { // odd length
		start--
		end++
		count++
	}
	for start >= 0 && end < n && s[start] == s[end] {
		count++
		start--
		end++
	}
	return count
}

func countSubstringswithdp(s string) int {
	chrs := []rune(s)
	n := len(chrs)

	dp := make([][]bool, len(chrs))
	for i := 0; i < len(chrs); i++ {
		dp[i] = make([]bool, len(chrs))
	}

	// single chars are all palindrome
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	for l := 2; l <= n; l++ {
		for st := 0; st < n; st++ {
			end := st + l - 1
			if end >= n {
				continue
			}
			if chrs[st] != chrs[end] {
				dp[st][end] = false
				continue
			}
			dp[st][end] = true // assume true first
			stn := st + 1      // start + 1
			endp := end - 1    // end - 1

			// reject all invalid stuff.
			if stn >= n || endp < 0 || stn > endp {
				continue
			}

			if !dp[stn][endp] {
				// fmt.Printf("marked false st: %d, en: %d, len: %d\n", st, end, l)
				dp[st][end] = false
			}

		}
	}

	// check in dp which all are marked true??
	total := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if dp[i][j] {
				total++
			}
		}
	}

	return total
}

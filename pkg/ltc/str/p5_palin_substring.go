package str

func longestPalindrome(s string) string {
	// return longestPalindromeRec(s)
	return longestPalinIter(s)
}


// approach
// DP[i][j] is true if s[i:j] is palindrome
// find palindromes of size 1, 2, 3 and so on, 
// making use of prev. size result to efficiently find next size results.
func longestPalinIter(s string) string {

	A := []rune(s)
	n := len(A)

	dp := make([][]bool, 0)
	for i := 0; i < n; i++ {
		dp = append(dp, make([]bool, n))
	}

	// single char
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	mx := 1
	mi = 0
	mj := 0

	// 2 char
	for i := 0; i < n-1; i++ {
		if A[i] == A[i+1] {
			mx = 2
			mi = i
			mj = i + 1
			dp[i][i+1] = true
		}
	}

	for sz := 1; sz <= n; sz++ {
		for st := 0; st < n; st++ {
			en := st + sz - 1
			// check if str[st...en] is a palindrome
			if st+1 < n && en < n && en > 1 && dp[st+1][en-1] && A[st] == A[en] {
				dp[st][en] = true
				if sz > mx {
					mx = sz
					mi = st
					mj = en
				}
			}
		}
	}

	return string(A[mi : mj+1])
}

// TODO - pretty messy, clean it up.
func longestPalindromeRec(s string) string {
	// a0, a1, a2, ... , an-2, an-1
	// f(i, j) = if (A[i] == A[j]) { if f(i+1, j-1) >= 0, then f(i+1, j-1) + 1, else 0)
	//  else return 0
	mi = -1
	mj = -1
	mv = -1
	A := []rune(s)
	n := len(A)
	dp := make([][]int, 0)

	for i := 0; i < n; i++ {
		dp = append(dp, make([]int, n))
	}

	// fill dp
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = -2
		}
	}

	f(0, len(A)-1, A, dp)
	// fmt.Printf("max is %d\n", mv)
	if mv <= 0 {
		return ""
	}
	return string(A[mi : mj+1])
}

var mi, mj, mv int

// returns length of palindrome starting at i, ending at j
func f(i, j int, A []rune, dp [][]int) int {
	// fmt.Printf("inside %d, %d\n", i, j)

	if i < 0 || i >= len(A) {
		return -1
	}
	if j < 0 || j >= len(A) {
		return -1
	}
	if dp[i][j] != -2 {
		return dp[i][j]
	}

	if i == j {
		if 1 > mv {
			mv = 1
			mi = i
			mj = j
			// fmt.Printf("got max equal case: i: %d j: %d max: %d\n", i, j, mv )
		}
		dp[i][j] = 1
		return 1
	}
	if i > j {
		dp[i][j] = 0
		return 0
	}
	// xSTRINGx, if STRING is palin
	// x

	mx := -1

	if A[i] == A[j] {
		// fmt.Printf("inside equal: %d, %d, %c, %c\n", i, j, A[i], A[j])
		next := f(i+1, j-1, A, dp)
		if next >= 0 {
			// mv = max(mv, next + 1)
			if next+2 > mv {
				mv = next + 2
				mi = i
				mj = j
				// fmt.Printf("got max: i: %d j: %d max: %d\n", i, j, mv )
			}
			mx = next + 2
			dp[i][j] = mx
			return mx
		}
	}

	// check for next solutions.
	f(i+1, j, A, dp)
	f(i, j-1, A, dp)
	// fmt.Printf("Exit: %d, %d, %d\n", i, j, mx)
	dp[i][j] = mx
	return mx
}

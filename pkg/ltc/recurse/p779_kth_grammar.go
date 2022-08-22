package recurse

func kthGrammar(n int, k int) int {
	// if k is even in the Nth stage, then A[k] = !A[k/2] from (n-1)th stage.
	// if k is odd in the Nth stage, then A[k] = A[k/2] from (n-1)th stage.
	if n == 1 {
		return 0
	}
	if k%2 == 1 {
		return kthGrammar(n-1, (k+1)/2)
	} else {
		val := kthGrammar(n-1, (k+1)/2)
		if val == 0 {
			return 1
		}
		return 0
	}
}

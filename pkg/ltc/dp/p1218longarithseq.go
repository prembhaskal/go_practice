package dp

func longestSubsequence(arr []int, difference int) int {
	return longestSubsequenceMap(arr, difference)
}
func longestSubsequenceMap(arr []int, difference int) int {
	// for each num, check if (num-diff) exists, if yes, what length ended at that.
	// use map for backward searching

	idxmap := make(map[int]int) // key = num, value = max length of valid index ending at num.

	// base cases
	idxmap[arr[0]] = 1
	maxlen := 1
	for i := 1; i < len(arr); i++ {
		need := arr[i] - difference
		currlen := 1
		prevlen, ok := idxmap[need]
		if ok {
			currlen = prevlen + 1
		}
		idxmap[arr[i]] = currlen

		maxlen = max1218(maxlen, currlen)
	}

	return maxlen
}

func max1218(a, b int) int {
	if a > b {
		return a
	}
	return b
}

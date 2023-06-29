package dp

func findNumberOfLIS(nums []int) int {
	lendp := make([]int, len(nums)) // lendp[i] = length of LIS ending at index 'i '
	cnt := make([]int, len(nums))   // cnt[i]   = count of LIS ending at index 'i'

	mxlen := 1
	for curr := 0; curr < len(nums); curr++ {
		// base case
		lendp[curr] = 1
		cnt[curr] = 1
		for prev := 0; prev < curr; prev++ {
			if nums[curr] > nums[prev] {
				if lendp[curr] < lendp[prev]+1 { // using prev, results in a longer sequence than seen till now.
					lendp[curr] = lendp[prev] + 1
					cnt[curr] = cnt[prev]
				} else if lendp[curr] == lendp[prev]+1 { // using prev, results in same seq as longest seen till now
					cnt[curr] += cnt[prev]
				}
			}
		}

		if lendp[curr] > mxlen {
			mxlen = lendp[curr]
		}
	}
	// count list
	totalcnt := 0
	for i := 0; i < len(nums); i++ {
		if lendp[i] == mxlen {
			totalcnt += cnt[i]
		}
	}

	return totalcnt
}

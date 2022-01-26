package binsearch

// https://leetcode.com/problems/first-bad-version/submissions/
func firstBadVersion(n int, isBadVersion func(int) bool) int {
	start := 1
	end := n
	for start <= end {
		mid := start + (end-start)/2
		if isBadVersion(mid) {
			if !isBadVersion(mid - 1) { // if previous is not bad
				return mid
			} else {
				end = mid - 1
			}
		} else {
			start = mid + 1
		}
	}

	return -1
}

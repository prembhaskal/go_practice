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

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

var isBadVersion func(int) bool

func firstBadVersion1(n int) int {
	start := 1
	end := n
	for start < end {
		mid := start + (end-start)/2
		// fmt.Printf("st: %d, mid: %d, end: %d\n", start , mid, end)
		midStatus := isBadVersion(mid)
		// mid1Status := isBadVersion(mid+1) // mid + 1  exists since start < end
		// if !midStatus && mid1Status {
		//     return mid + 1
		// }
		if !midStatus {
			start = mid + 1
		} else {
			end = mid
		}
	}

	if isBadVersion(start) { // at end start = end.
		return start
	}

	return -1
}

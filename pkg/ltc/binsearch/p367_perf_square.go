package binsearch

func isPerfectSquare(num int) bool {
	start := 1
	end := num
	for start <= end {
		mid := start + (end-start)/2
		// fmt.Printf("start: %d, mid: %d, end: %d\n", start, mid, end)
		quot := num / mid
		rem := num % mid
		if mid == quot && rem == 0 {
			return true
		} else if mid > num/mid {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return false
}

func isPerfectSquare1(num int) bool {
	start := 1
	end := num
	for start <= end {
		mid := start + (end-start)/2
		// fmt.Printf("start: %d, mid: %d, end: %d\n", start, mid, end)
		if mid == num/mid && mid*mid == num {
			return true
		} else if mid > num/mid {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return false
}

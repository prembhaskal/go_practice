package binsearch

func mySqrt(x int) int {
	return mySqrt1(x)
}

// search between 1 to x
func mySqrt1(x int) int {
	start := 1
	end := x
	for start <= end {
		mid := start + (end-start)/2
		sqr := mid * mid
		if sqr == x {
			return mid
		} else if sqr < x {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return end
}

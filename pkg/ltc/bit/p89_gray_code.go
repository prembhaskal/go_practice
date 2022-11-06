package bit

var total int

func grayCode(n int) []int {
	total = 1 << n
	gray := make(map[int]bool)
	gray[0] = true
	grays := []int{0}
	return calGrayCode(n, 0, gray, grays)
}

func calGrayCode(n, num int, gray map[int]bool, grays []int) []int {
	// if we reached here, where we get 1<<n elements.
	if len(grays) == total {
		// fmt.Printf("done: cnt: %d, total: %d\n", cnt, total)
		return grays
	}

	// try flipping 1 bit at every position and recurse from there.
	for i := 0; i < n; i++ {
		// flip ith bit
		next := flipBit(num, i)
		if gray[next] {
			continue
		}
		gray[next] = true
		grays = append(grays, next)

		finalgray := calGrayCode(n, next, gray, grays)
		if finalgray != nil {
			return finalgray
		}
		grays = grays[:len(grays)-1]
		gray[next] = false
	}
	return nil
}

func flipBit(num int, idx int) int {
	mask := 1 << idx
	return num ^ mask
}

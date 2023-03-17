package dp

func minimumDifference(nums []int) int {
	n1 := 0
	s1 := 0
	n2 := 0
	s2 := 0
	return f2035(nums, 0, n1, s1, n2, s2)
}

var inf int = 1000000000 // 10 ^ 9

// TODO - time limit exceed, optimize
func f2035(nums []int, curr, n1, s1, n2, s2 int) int {
	// fmt.Printf("debug, s1: %d, n1: %d, s2: %d, n2: %d, curr: %d\n", s1, n1, s2, n2, curr)
	if curr == len(nums) {
		// check
		if n1 != n2 || n1 != len(nums)/2 {
			return inf
		}

		// fmt.Printf("end, s1: %d, s2: %d, len(s1):%d, len(s2):%d\n", s1, s2, n1, n2)

		return abs2035(s1 - s2)
	}

	// nums[curr]
	// add in set1
	val1 := f2035(nums, curr+1, n1+1, s1+nums[curr], n2, s2)
	// add in set2
	val2 := f2035(nums, curr+1, n1, s1, n2+1, s2+nums[curr])
	return min2035(val1, val2)
}

func min2035(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs2035(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

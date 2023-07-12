package slidewin

func findMaxConsecutiveOnes(nums []int) int {
	// return findMaxConsecutiveOnesSlideWindow(nums)
	return findMaxConsecutiveOnesFollowup(nums)
}

// left .... right , maintain varaint that we have at max 1 zero in between
// reduce left incr. by remembering index of last zero.
func findMaxConsecutiveOnesFollowup(nums []int) int {
	left := 0
	right := 0

	numzero := 0
	maxlen := 0
	lastzero := -1

	for ; right < len(nums); right++ {
		if nums[right] == 0 {
			numzero++
		}

		if numzero == 2 {
			left = lastzero + 1
			numzero = 1
		}

		maxlen = max487(maxlen, right-left+1)

		if nums[right] == 0 {
			lastzero = right
		}
	}

	return maxlen
}

// left .... right , maintain varaint that we have at max 1 zero in between
func findMaxConsecutiveOnesSlideWindow(nums []int) int {
	left := 0
	right := 0

	numzero := 0
	maxlen := 0

	for ; right < len(nums); right++ {
		if nums[right] == 0 {
			numzero++
		}

		if numzero == 2 {
			// move left until we have only 1 zero
			for nums[left] != 0 {
				left++
			}
			numzero--
			left++
		}
		maxlen = max487(maxlen, right-left+1)

	}

	return maxlen
}

func findMaxConsecutiveOnesPrefixSuffix(nums []int) int {
	// try with prefix sum
	fwd := make([]int, len(nums))

	maxones := 0

	currsum := 0
	for idx, v := range nums {
		if v == 0 {
			currsum = 0
		} else {
			currsum++
		}
		fwd[idx] = currsum
		maxones = max487(maxones, currsum)
	}

	// fmt.Printf("fwd: %+v\n", fwd)

	rev := make([]int, len(nums))
	currsum = 0
	for idx := len(nums) - 1; idx >= 0; idx-- {
		if nums[idx] == 0 {
			currsum = 0
		} else {
			currsum++
		}
		rev[idx] = currsum
	}

	// fmt.Printf("rev: %+v\n", rev)

	for idx, v := range nums {
		if v == 0 {
			onleft := 0
			onright := 0
			if idx > 0 {
				onleft = fwd[idx-1]
			}
			if idx+1 < len(nums) {
				onright = rev[idx+1]
			}

			maxones = max487(maxones, onleft+onright+1)
		}
	}

	return maxones
}

func max487(a, b int) int {
	if a > b {
		return a
	}
	return b
}

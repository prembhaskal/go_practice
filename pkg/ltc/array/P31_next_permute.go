package array

func nextPermutation(nums []int) {
	// a1, a2, a3, a4, ..., ai, aj, .... an
	// search from last find (ai, aj) pair such that ai < aj

	n := len(nums)
	idx := -1
	for i := n - 2; i >= 0; i-- {
		ai := nums[i]
		aj := nums[i+1]
		if ai < aj {
			idx = i
			break
		}
	}

	// fmt.Printf("swipe index is %d\n", idx)
	if idx != -1 {
		// place to swipe is idx
		// search from last, min(ax) such that ax > ai
		imin := -1
		for i := n - 1; i > idx; i-- {
			if nums[i] > nums[idx] {
				if imin == -1 {
					imin = i
				} else {
					if nums[i] < nums[imin] {
						imin = i
					}
				}
			}
		}

		// fmt.Printf("imin: %d\n", imin)

		nums[idx], nums[imin] = nums[imin], nums[idx]
	}

	// ideally, we need to sort from idx+1 to end
	// already reverse sort, just write it out in reverse
	j := n - 1
	for i := idx + 1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]

		i++
		j--
	}
}

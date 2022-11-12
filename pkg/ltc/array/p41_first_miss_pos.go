package array

// TODO optimize this, only 2 loops should be enough, ignore >= 0
func firstMissingPositive(nums []int) int {
	lastptr := len(nums) - 1
	for ; lastptr >= 0; lastptr-- {
		// fmt.Printf("lastptr: %d, nums[lastptr]: %d\n", lastptr, nums[lastptr])
		if nums[lastptr] > 0 {
			break
		}
	}

	// fmt.Printf("1. lastptr: %d\n", lastptr)

	if lastptr < 0 {
		return 1
	}

	for i := 0; i < lastptr && lastptr >= 0; {
		if nums[i] <= 0 {
			nums[lastptr], nums[i] = nums[i], nums[lastptr]
			lastptr--
		} else {
			i++
		}
	}

	if nums[lastptr] <= 0 {
		lastptr--
	}

	// fmt.Printf("2. lastptr: %d\n", lastptr)

	// fmt.Printf("left +ves, right non +ve: %v\n", nums)
	// all positives on left side. negative on right side.
	// lastptr is last positive num
	// 2 1 1 0 -1 -1

	for i := 0; i <= lastptr; {
		if nums[i]-1 != i {
			// fmt.Printf("last loop: nums[i]:%d, i:%d\n", nums[i], i)
			if nums[i]-1 > lastptr { // we cannot swap this
				i++
				continue
			}
			idx := nums[i] - 1
			// check if the dest. already has right value, then dont' swap
			if nums[idx] == idx+1 {
				i++
				continue
			}

			// swap i and nums[i]-1
			nums[i], nums[idx] = nums[idx], nums[i]
			continue
		} else {
			i++
		}
	}

	i := 0
	for ; i <= lastptr; i++ {
		if i != nums[i]-1 {
			return i + 1
		}
	}

	return i + 1
}

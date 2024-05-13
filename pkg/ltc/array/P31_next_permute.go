package array

import "slices"

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

// bit simpler code
func nextPermutation1(nums []int) {
	// 1, 3, 4, 2, 5, 7, 6
	// 1, 3, 4, 2, 6, 5, 7

	// 5, 2, 4, 3, 1
	// search element from last to first, not in sorted order
	// exchange with smallest element bigger than it from right side
	// sort the right side.
	// 5, 3, 1, 2, 4

	// 4, 3, 2, 1
	// skip exchange part, just sort the right side
	// 1, 2, 3, 4

	// find element
	n := len(nums)
	exch_idx := -1
	for i := n - 2; i >= 0; i-- {
		curr := nums[i]
		prev := nums[i+1]
		if curr < prev {
			exch_idx = i
			break
		}
	}

	// fmt.Printf("Exch idx is %d\n", exch_idx)
	if exch_idx != -1 {
		// find element to swap and swap, this search could be made binary search too, since in sorted form.
		for i := n - 1; i > exch_idx; i-- {
			if nums[i] > nums[exch_idx] {
				// swap
				nums[i], nums[exch_idx] = nums[exch_idx], nums[i]
				break
			}
		}
	}

	// tosort := nums[exch_idx+1:]
	// sort.Ints(tosort)

	// already reverse sorted.
	rev(nums, exch_idx+1, n-1)
}

func rev(a []int, s, e int) {
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

}

func nextPermutation2(nums []int) {
	n := len(nums)
	for i := n - 2; i >= 0; i-- {
		// i i+1
		if nums[i] < nums[i+1] {
			// i is anomaly

			// find element to swap from back (since already sorted)
			m := -1
			for j := len(nums) - 1; j >= i+1; j-- {
				if nums[j] > nums[i] {
					m = j
					break
				}
			}

			nums[i], nums[m] = nums[m], nums[i]
			// reverse remaining ones (already in descending order)
			slices.Reverse(nums[i+1:])
			return
		}
	}

	slices.Reverse(nums)
}

package binsearch

func singleNonDuplicate(nums []int) int {
	start := 0
	end := len(nums) - 1

	// double means even(i) and odd(i+1) should have same number

	for start < end {
		mid := start + (end-start)/2

		// turns out below check is not needed at all, the algo nicely converges to single number
		// check if mid is the one which appears once?
		// if mid > 0 {
		//     if nums[mid-1] != nums[mid] && nums[mid] != nums[mid+1] {
		//         return nums[mid]
		//     }
		// } else {
		//     if nums[0] != nums[1] {
		//         return nums[0]
		//     }
		// }

		if mid%2 == 0 {
			if nums[mid] == nums[mid+1] {
				start = mid + 2
			} else {
				end = mid
			}
		} else {
			if nums[mid-1] == nums[mid] {
				start = mid + 1
			} else {
				end = mid // mid-1. if mid is the one, it would be already caught above, so safe to ignore it.
			}
		}
	}

	return nums[start]

	// 0 1 2 3 4 5 6
	// 1 1 2 3 3 4 4
	// start = 0
	// end = 6
	// mid = 3
	// start = 0, end = 2 mid = 1
	//

}

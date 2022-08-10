package array

func lengthOfLIS(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	// init tail table, where index = len of subsequence.
	tt := make([]int, 2)

	// DP[i] = {max{DP[j]]} where j = i-1 to 0 and A[j] < A[i]}
	// we use tail table to search in the smallest element larger than current element.

	// initial condition
	tt[1] = nums[0]

	for i := 1; i < len(nums); i++ {
		last_elem := tt[len(tt)-1]
		if nums[i] > last_elem {
			tt = append(tt, nums[i])
			continue
		}

		// else search element in tail table
		elem_idx := binSearchCeilIndex(tt, nums[i])
		tt[elem_idx] = nums[i]

		// fmt.Printf("nums[%d]=%d, elem_idx=%d\n", i, nums[i], elem_idx)
	}

	return len(tt) - 1
}

// if number == key, return its index.
// else search smallest number > key and return its index.
func binSearchCeilIndex(tt []int, key int) int {
	low := 1
	high := len(tt) - 1
	for low < high {
		mid := (low + high) / 2
		// if tt[mid] == key {
		//     return mid
		// }
		if tt[mid] < key {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return high
}

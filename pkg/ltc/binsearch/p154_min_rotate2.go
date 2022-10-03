package binsearch

func findMin2(nums []int) int {
	//     4, 4, 4, 4, 4, 4, 1, 4, 4
	//     6, 7, 8, 9, 1, 2, 3, 4, 5
	//     1, 2, 3
	//     3, 4, 2
	//     3, 1, 2
	//     4, 5
	//     5, 4

	start := 0
	end := len(nums) - 1
	for start < end {
		mid := start + (end-start)/2
		// fmt.Printf("start: %d, mid: %d, end: %d\n", start, mid, end)
		// if nums[mid] > nums[end] { // right not sorted, elem lies here
		//     start = mid + 1
		// } else if nums[start] > nums[mid] { // left not sorted
		//     end = mid
		// } else if nums[start] < nums[mid] { // left sorted
		//     end = mid
		// } else if nums[mid] == nums[end] {// both are equal, we cannot choose a side, reduce end by 1 to remove 1 duplicate at least.
		//     end--
		// } else if nums[start] == nums[mid] { // both are equal, we cannot choose a side, reduce end by 1 to remove 1 duplicate at least.
		//     end--
		// }
		// same as above but conditions merged
		if nums[mid] > nums[end] {
			start = mid + 1
		} else if nums[mid] == nums[end] {
			end--
		} else {
			end = mid
		}
	}

	// start had min at end of loop.
	return nums[start]
}

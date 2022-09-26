package binsearch

func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1

	for start < end {
		mid := start + (end-start)/2
		// fmt.Printf("start: %d, mid: %d, end: %d\n", start, mid , end)
		// check if right is not sorted. then min lies here
		if nums[mid] > nums[end] {
			start = mid + 1 // mid + 1 because mid is converging to left side, take 2 nums as example and check.
			// } else if nums[start] == nums[mid] { // 2 elements are left., MERGED into below case.
			//     end = mid
		} else {
			end = mid
		}
	}

	return nums[start]
}

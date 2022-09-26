package binsearch

func findPeakElement(nums []int) int {
	// start , mid , end
	start := 0
	// n := len(nums)
	end := len(nums) - 1

	for start < end {
		mid := start + (end-start)/2
		// fmt.Printf("start: %d, end: %d, mid: %d\n", start, end, mid)

		// boundary case start == mid < end, eg . [3, 4]  or [4, 3]
		// peak towards right, since slope rising
		if nums[mid] < nums[mid+1] { // mid + 1 exists because sart < end and mid truncates to lower of start,end
			start = mid + 1
		} else { // mid-1 > mid , peak towards left
			end = mid
		}

	}
	return start
}

func findPeakElement1(nums []int) int {
	// start , mid , end
	start := 0
	n := len(nums)
	end := len(nums) - 1

	for start < end {
		mid := start + (end-start)/2
		// fmt.Printf("start: %d, end: %d, mid: %d\n", start, end, mid)

		// we have at least 3 elements.
		if mid-1 >= 0 && mid+1 <= n-1 {
			// check if this is peak
			if nums[mid-1] < nums[mid] {
				if nums[mid] > nums[mid+1] {
					return mid
				}
				// peak towards right, since slope rising
				if nums[mid] < nums[mid+1] {
					start = mid + 1
				}
			} else { // mid-1 > mid , peak towards left
				end = mid
			}
		} else {
			break // with less than 2 elements, we just return whichever is maximum.
		}

	}
	if start == end {
		return start
	}
	if nums[start] > nums[end] {
		return start
	}

	return end
}

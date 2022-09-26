package binsearch

func search(nums []int, target int) int {
	return search2(nums, target)
}

//  HINT for better understand write all conditions (including start==mid || mid == end) first separately
// and then merge the common conditions to arrive at below
func search1(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	key := target

	for start <= end {
		// fmt.Printf("start: %d, end: %d\n", start, end)
		mid := start + (end-start)/2
		// fmt.Printf("mid: %d, a[mid]=%d\n", mid, nums[mid])
		if nums[mid] < key {
			if nums[mid] < nums[end] && nums[end] < key {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else if nums[mid] > key {
			if nums[start] <= nums[mid] && nums[start] > key {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			return mid
		}
	}
	return -1
}

func search2(nums []int, target int) int {
	var start, end, mid int
	start = 0
	end = len(nums) - 1

	for start <= end {
		mid = start + (end-start)/2

		if nums[mid] == target {
			return mid
		}

		// check if left part is sorted
		if nums[start] <= nums[mid] {
			// if yes, check if our target lies here
			if target >= nums[start] && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else if nums[mid] <= nums[end] { // check if right part is sorted.
			if target > nums[mid] && target <= nums[end] { // check if target lies in the sorted area
				start = mid + 1
			} else {
				end = mid - 1
			}
		}

	}

	return -1
}

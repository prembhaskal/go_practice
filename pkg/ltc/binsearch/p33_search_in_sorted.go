package binsearch

func search(nums []int, target int) int {
	return search2(nums, target)
}

//	HINT for better understand write all conditions (including start==mid || mid == end) first separately
//
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

// one more approach
// check if array part rotated
// if yes, check in sorted part of the rotated array
// if lies in sorted part, check in that part, else check in other part.
// if not rotated, do normal bin search
func search3(nums []int, target int) int {
	n := len(nums)

	start := 0
	end := n - 1

	// 0 1 2 3 4 5 6 7
	// 6 7 0 1 2 3 4 5
	// 3 4 5 6 7 0 1 2

	//  start .... mid .... end
	//         t
	//                   t

	for start < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}

		if nums[start] > nums[end] { // rotated.
			if nums[mid] < nums[end] { // right side sorted
				if target > nums[mid] && target <= nums[end] { // number lies in sorted part
					start = mid + 1
				} else {
					end = mid - 1 // if it is not in right side, it will be left side
				}
			} else { // left side sorted
				if target >= nums[start] && target < nums[mid] { // number lies in sorted part
					end = mid - 1
				} else {
					start = mid + 1
				}
			}

		} else { // not rotated, do usual search
			if target > nums[mid] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	// boundary case
	// case 1
	// 3 4
	// 6 7
	// mid = 3 nums[mid] = 6
	// case 2
	// 3 4
	// 7 2
	// mid = 3
	if nums[start] == target {
		return start
	}
	return -1
}

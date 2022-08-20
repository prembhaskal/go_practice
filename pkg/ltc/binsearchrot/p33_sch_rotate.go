package binsearchrot

// [9, 12, 15, -2, -1 , 0]
// search solves this problem https://leetcode.com/problems/search-in-rotated-sorted-array/submissions/
func search(nums []int, target int) int {
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
				start = mid + 1 // if not, then we get rid of it and reduce problem area.
			}
		} else if nums[mid] <= nums[end] { // check if right part is sorted.
			if target > nums[mid] && target <= nums[end] { // check if target lies in the sorted area
				start = mid + 1
			} else {
				end = mid - 1 // if not, then we get rid of it and reduce problem area.
			}
		}

	}

	return -1
}

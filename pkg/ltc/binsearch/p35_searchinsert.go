package binsearch

// https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			if mid > 0 && target > nums[mid-1] {
				return mid // mid is the index where it can be inserted.
			}
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	// if target is out of range of given array
	if target < nums[0] {
		return 0
	} else {
		return len(nums)
	}
}

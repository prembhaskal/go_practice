package binsearch

import "fmt"

func searchdup2(nums []int, target int) bool {

	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return true
		}
		// special case
		if nums[low] == nums[mid] && nums[mid] == nums[high] {
			return linsearch(nums[low:high+1], target)
			// low++
			// high--
			// continue
		}

		// boundary cases
		if low == mid { // only 2 elements left, mid is already checked above.
			low = mid + 1
		} else if nums[low] <= nums[mid] { // check if left is in sorted order
			if target >= nums[low] && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else { // right is sorted
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		fmt.Printf("low: %d, mid: %d, high: %d\n", low, mid, high)
	}
	return false
}

func linsearch(nums []int, tgt int) bool {
	for _, v := range nums {
		if v == tgt {
			return true
		}
	}
	return false
}

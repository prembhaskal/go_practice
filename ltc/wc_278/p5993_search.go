package wc_278

import "sort"

func findFinalValue(nums []int, original int) int {
    sort.Ints(nums)

	max := nums[len(nums)-1]

	for original <= max {
		found := search(nums, original)
		if found != -1 {
			original = original * 2
		} else {
			break
		}
	}

	return original
}

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (start + end) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}

	}

	return -1
}
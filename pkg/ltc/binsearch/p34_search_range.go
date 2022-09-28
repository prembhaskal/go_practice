package binsearch

import "fmt"

func searchRange(nums []int, target int) []int {
	low := getlowerbound(nums, target)
	if low == -1 {
		return []int{-1, -1}
	}
	high := getupperbound(nums, target)

	return []int{low, high}
}

func getupperbound(ar []int, key int) int {
	start := 0
	end := len(ar) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		fmt.Printf("start: %d, mid: %d, end: %d\n", start, mid, end)
		if ar[mid] < key {
			start = mid + 1
		} else if ar[mid] > key {
			end = mid
		} else if ar[mid] == key {
			start = mid
		}
	}

	if ar[end] == key {
		return end
	}
	if ar[start] == key {
		return start
	}
	return -1
}

func getlowerbound(ar []int, key int) int {
	if len(ar) == 0 {
		return -1
	}
	start := 0
	end := len(ar) - 1

	for start < end {
		mid := start + (end-start)/2
		if ar[mid] < key {
			start = mid + 1
		} else if ar[mid] > key {
			end = mid
		} else if ar[mid] == key {
			end = mid
		}
	}

	if ar[end] != key {
		return -1
	}

	return end
}

func searchRange1(nums []int, target int) []int {
	// 1 , 2 , 3, 3, 3, 3, 4 ,5

	start := 0
	end := len(nums) - 1

	if end == -1 {
		return []int{-1, -1}
	}

	// find left index. num[idx] == target && (num[idx] != target || idx == 0 )
	leftidx := -1
	for start+1 < end {
		mid := start + (end-start)/2
		// fmt.Printf("start: %d, mid: %d, end: %d\n", start, mid, end)
		if nums[mid] == target && nums[mid-1] != target {
			leftidx = mid
			break
		} else if nums[mid] >= target {
			end = mid
			// } else if nums[mid] == target {
			//     end = mid
		} else if nums[mid] < target {
			start = mid + 1
		}
	}

	// fmt.Printf("start: %d, end: %d\n", start, end)

	if leftidx == -1 {
		if nums[start] == target {
			leftidx = start
		} else if nums[end] == target {
			leftidx = end
		} else {
			return []int{-1, -1}
		}
	}
	// fmt.Printf("leftidx: %d\n", leftidx)

	// find right index, num[idx] == target && ( num[idx] != target || idx == len(nums)-1)
	rightidx := -1
	start = leftidx
	end = len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target && nums[mid+1] != target {
			rightidx = mid
			break
		} else if nums[mid] > target {
			end = mid
			// } else if nums[mid] == target {
			//     start = mid
		} else if nums[mid] <= target {
			start = mid
		}
	}
	if rightidx == -1 {
		if nums[end] == target {
			rightidx = end
		} else if nums[start] == target {
			rightidx = start
		}
	}
	return []int{leftidx, rightidx}
}

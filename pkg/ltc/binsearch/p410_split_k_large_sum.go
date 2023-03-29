package binsearch

import "fmt"

func splitArray(nums []int, m int) int {
	// return 0

	fmt.Printf("\n********** new testcase nums: %v **************\n\n", nums)
	total := 0
	for _, v := range nums {
		total = total + v
	}

	fmt.Printf("total:%d\n", total)

	start := 0
	end := total

	for start < end {
		mid := start + (end-start)/2
		fmt.Printf("start: %d, mid: %d, end: %d\n", start, mid, end)
		// if for mid, if m or less split possible, then for num > mid, it is possible to split less than m
		if checkMinSplitPossible(nums, mid, m) <= m {
			fmt.Printf("%d-split possible with mid:%d\n", m, mid)
			end = mid
			// fmt.Printf("end is set to %d\n", end)
		} else {
			// if for mid, if m splits not possible, then num < mid, it is not possible to split <=m, all splits will be more than m.
			fmt.Println("m-split not possible")
			start = mid + 1
		}
	}

	return start
}

// check the minimum split possible, to achieve the sum.
// for eg. [2, 2, 2, 2, 2, 2, 2, 2, 2, 2] -- 10 items
// for max_sum=4, min split possible is 5 times (each having 2 items), we cannot have fewer splits.
//
//	for max_sum = 3 or less, min split possible would be higher than 5.
//
// for max_sum=8, min split possible is 3 times ([2,2,2,2] [2,2,2,2] [2,2])
//
//	for max_sum >= 9, min split possible is at max 3 times.
func checkMinSplitPossible(nums []int, largesum, m int) int {
	splits := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > largesum {
			return m + 1 // we cannot split this array.
		}
		sum = sum + nums[i]
		if sum > largesum {
			splits++
			sum = nums[i]
		}
		fmt.Printf("i: %d, splits: %d\n", i, splits)
	}
	fmt.Printf("splits are %d\n", splits+1)
	return splits + 1 // +1 because last one won't get in the if condition.
}

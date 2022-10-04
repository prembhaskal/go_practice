package binsearch

import (
	"fmt"
	"sort"
)

// sort the nums
// pairwise difference
// pair of diff 1 -- n-1
// pair of diff 2 -- n-2

// order (n-1) (n-2) (n-3) ... (1)

// reverse (1) (2) (3) ... (n-1)th
func smallestDistancePair(nums []int, k int) int {
	// binary search where the kth part lies
	n := len(nums)

	key := (n * (n - 1) / 2) - k + 1
	fmt.Printf("key is %d\n", key)
	// key := k

	start := 1
	end := n
	for start < end {
		mid := start + (end-start)/2
		bstart := mid*(mid-1)/2 + 1 // batch start
		bend := mid * (mid + 1) / 2 // batch end

		// fmt.Printf("start: %d, end: %d, mid: %d, bstart: %d, bend: %d\n", start, end, mid, bstart, bend)
		// if key >= bstart && key <= bend {
		// 	break
		// } else
		if key < bstart {
			end = mid
		} else if key > bend {
			start = mid + 1
		} else {
			start = mid
			break
		}
	}

	// for start = 1; start <= end; start++ {
	// 	bstart := start * (start - 1) / 2 + 1
	// 	bend := start * (start + 1) / 2
	// 	fmt.Printf("start: %d, end: %d, bstart: %d, bend: %d\n", start, end, bstart, bend)
	// 	if key >= bstart && key <= bend {
	// 		break
	// 	}
	// }

	// start has min.
	// return n - start
	fmt.Printf("found start batch: %d\n", start)
	diff := n - start
	fmt.Printf("diff batch start: %d\n", diff)
	// sort.Ints(nums)
	sort.Sort(sort.IntSlice(nums))
	fmt.Printf("sorted nums: %v\n", nums)
	currbatch := make([]int, 0)
	for i := 0; i+diff < n; i++ {
		currbatch = append(currbatch, abs(nums[i]-nums[i+diff]))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(currbatch)))

	fmt.Printf("curr batch is %v\n", currbatch)

	gap := key - (start * (start - 1) / 2 + 1)
	fmt.Printf("gap is %d\n", gap)
	return currbatch[gap]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

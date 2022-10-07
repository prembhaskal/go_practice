package binsearch

import (
	"fmt"
	"sort"
)

// sort the nums
// define paircount(m) =  where count_of_pairs_with_dist <= m
// possible dist = [0, max(nums)-min(nums)]
// binary search for the dist, for which paircount(dist-1) < k && paircount(dist+1) > k
func smallestDistancePair(nums []int, k int) int {
	sort.Sort(sort.IntSlice(nums))
	fmt.Printf("sorted array: %v\n", nums)

	n := len(nums)

	start := 0
	end := nums[n-1] - nums[0]

	for start < end {
		mid := start + (end-start)/2
		cnt := paircount(nums, mid)
		if cnt < k {
			start = mid + 1
			// } else if cnt == k {
			// 	end = mid
		} else if cnt >= k {
			end = mid
		}
	}

	return start
}

// needs arr to be sorted
// paircount(m) =  where count_of_pairs_with_dist <= m
func paircount(arr []int, dist int) int {
	// 1, 3, 4, 5, 6
	count := 0
	n := len(arr)
	i := 0
	j := 1
	for ; i < n; i++ {
		for ; j < n; j++ {
			if arr[j]-arr[i] <= dist {
				count = count + (j - i)
			} else {
				break
			}
		}
	}
	return count
}

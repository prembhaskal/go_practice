package sort

import (
	// "fmt"
	"math/rand"
)

func sortArray(nums []int) []int {
	// quicksort(nums, 0, len(nums)-1)
	quicksortrandom(nums, 0, len(nums)-1)
	return nums
}

func quicksortrandom(nums []int, start, end int) {
	// fmt.Printf("quicksort start: %d, end: %d, nums: %v\n", start, end, nums)
	if start >= end {
		return
	}

	// find random number between start and end & swap with end element.
	n := rand.Intn(end-start+1) + start
	nums[end], nums[n] = nums[n], nums[end]

	mid := partition(nums, start, end)
	// fmt.Printf("pivot index: %d\n", mid)
	quicksortrandom(nums, start, mid-1)
	quicksortrandom(nums, mid+1, end)
}

func quicksort(nums []int, start, end int) {
	// fmt.Printf("quicksort start: %d, end: %d, nums: %v\n", start, end, nums)
	if start >= end {
		return
	}
	// pivot := end
	mid := partition(nums, start, end)
	// fmt.Printf("pivot index: %d\n", mid)
	quicksort(nums, start, mid-1)
	quicksort(nums, mid+1, end)
}

///  start ... i i+1 ... j ... end/pivot
// i last element which is less than x
// (i, j] more than x
// j last element yet to be checked.
// i+1 is where pivot will sit.

func partition(nums []int, start, end int) int {
	// fmt.Printf("start: %d, end: %d\n", start, end)
	key := nums[end] // last element is pivot
	// fmt.Printf("partition pivot idx: %d, value: %d\n", end, key)

	i := start - 1
	j := start
	for ; j < end; j++ {
		if nums[j] <= key {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[end] = key, nums[i+1]
	return i + 1
}

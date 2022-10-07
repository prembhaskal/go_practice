package binsearch

import "fmt"

func splitArray(nums []int, m int) int {
	return 0
	// total := 0
	// for _, v := range nums {
	// 	total = total + v
	// }

	// fmt.Printf("total:%d\n", total)

	// start := 0
	// end := total

	// for start < end {
	// 	mid := start + (end-start)/2
	// 	fmt.Printf("start: %d, mid: %d, end: %d\n", start, mid, end)
	// 	if isMSplitPossible(nums, mid, m) {
	// 		fmt.Println("m-split possible")
	// 		start = mid
	// 	} else {
	// 		fmt.Println("m-split not possible")
	// 		start = mid + 1
	// 	}
	// }

	// return end
}

// check the minimum split possible, to achieve the sum.
func checkMinSplitPossible(nums []int, largesum, m int) int {
	sublen := 0

	for i := 0; i < len(nums); i++ {
		sum := 0
		subarrcnt := 0
		for ; i < len(nums); i++ {
			sum = sum + nums[i]
			subarrcnt++
			// fmt.Printf("sum is %d, index i: %d\n", sum, i)
			if sum > largesum {
				sublen++
				subarrcnt--
				i--
				break
			}
		}
		if subarrcnt == 0 {
			fmt.Println("sub arr cnt is zero")
			return m+1
		}
	}
	fmt.Printf("sublen is %d\n", sublen)
	// if sublen >= m {
	// 	return true
	// }
	// return false
	return sublen
}

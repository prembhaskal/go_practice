package binsearch

func findDuplicate(nums []int) int {
	return findDupBinSearch(nums)
}

func findDupBinSearch(nums []int) int {
	// less_count of num = count of num which is less than or equal to it.
	// find smallest number, for less_count(num) is more than num itself.

	start := 1
	end := len(nums)

	for start < end {
		mid := start + (end-start)/2
		lesscnt := lesscount(mid, nums)
		if lesscnt < mid {
			start = mid + 1
		} else if lesscnt == mid {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return start
}

func lesscount(a int, nums []int) int {
	cnt := 0
	for _, v := range nums {
		if v <= a {
			cnt++
		}
	}
	return cnt
}

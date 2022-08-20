package wc_279

import (
	"sort"
)

// https://leetcode.com/contest/weekly-contest-279/problems/smallest-value-of-the-rearranged-number/
func smallestNumber(num int64) int64 {
	digits := make([]int, 0, 32)

	tmp := num
	if num < 0 {
		num = -num
	}

	for num > 0 {
		rem := int(num % 10)
		num = num / 10
		digits = append(digits, rem)
	}

	if tmp > 0 {
		sort.Sort(sort.IntSlice(digits))

		i := 0
		for ; i < 32; i++ {
			if digits[i] != 0 {
				break
			}
		}

		// swap 1st non zero
		digits[0], digits[i] = digits[i], digits[0]

		return getNum(digits)

	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(digits)))
		return -getNum(digits)

	}
}

func getNum(digits []int) int64 {
	var num int64
	for _, v := range digits {
		num = num * 10
		num = num + int64(v)
	}

	return num
}

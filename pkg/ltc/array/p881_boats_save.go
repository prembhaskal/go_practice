package array

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	// stable sort is quite slow btw, also it does change underlying slice
	// peopleSrt := sort.IntSlice(people)
	// sort.Stable(peopleSrt)

	st := 0
	end := len(people) - 1
	boats := 0
	for st <= end {
		if people[st]+people[end] <= limit {
			st++
		}
		boats++
		end-- // end is always used in all cases, question is only for start.
	}

	return boats
}

package dp

import (
	"sort"
)

func maxValue(events [][]int, k int) int {
	// return maxValueBruteForce(events, k)
	return maxValueDPSolve(events, k)
}

func maxValueBruteForce(events [][]int, k int) int {
	// sort the events, starttime, endtime
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] < events[j][0] {
			return true
		}
		if events[i][0] > events[j][0] {
			return false
		}
		return events[i][1] < events[j][1]
	})

	return maxValRec(events, 0, k)
}

// max value obtained from curr onward choosing 'k' events
func maxValRec(events [][]int, curr, k int) int {
	if k == 0 || curr >= len(events) { // end of the road
		return 0
	}

	// choose cur
	choose := 0
	// find next non-overlap events
	next := curr + 1
	for ; next < len(events); next++ {
		if events[next][0] > events[curr][1] {
			break
		}
	}
	choose = maxValRec(events, next, k-1) + events[curr][2]

	// not choose curr
	notchoose := maxValRec(events, curr+1, k)
	return max1751(choose, notchoose)
}

func maxValueDPSolve(events [][]int, k int) int {
	// sort the events, starttime, endtime
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] < events[j][0] {
			return true
		}
		if events[i][0] > events[j][0] {
			return false
		}
		return events[i][1] < events[j][1]
	})

	memo = make(map[pair1125]int)
	return maxValDP(events, 0, k)
}

var memo map[pair1125]int

func maxValDP(events [][]int, curr, k int) int {
	if k == 0 || curr >= len(events) { // end of the road
		return 0
	}

	currpair := pair1125{curr, k}
	if val, ok := memo[currpair]; ok {
		return val
	}

	// choose cur
	// find next non-overlap events using bin search upper bound, another option is to do linear search
	choose := 0
	// next := curr + 1
	// for ; next < len(events); next++ {
	//     if events[next][0] > events[curr][1] {
	//         break
	//     }
	// }
	next := upper_bound(events, curr+1, events[curr][1])
	if next == -1 {
		next = len(events)
	}

	choose = maxValDP(events, next, k-1) + events[curr][2]

	// not choose curr
	notchoose := maxValDP(events, curr+1, k)
	memo[currpair] = max1751(choose, notchoose)
	return memo[currpair]
}

func upper_bound(events [][]int, start int, bound int) int {
	low := start
	high := len(events) - 1
	ans := -1
	for low <= high {
		mid := low + (high-low)/2
		if events[mid][0] > bound {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}

type pair struct {
	a int
	b int
}

func max1751(a, b int) int {
	if a > b {
		return a
	}
	return b
}

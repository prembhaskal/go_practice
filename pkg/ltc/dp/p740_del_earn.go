package dp

import "sort"

func deleteAndEarn(nums []int) int {
	// return deleteAndEarnRec(nums)
	// return deleteAndEarnIter1(nums)
	return deleteAndEarnIter2(nums)
}

func deleteAndEarnIter1(nums []int) int {
	pick := make([]int, 10001)
	skip := make([]int, 10001)

	ar := make([]int, 10001)
	maxnum := -1
	for _, v := range nums {
		ar[v] = ar[v] + v // calculating earnings too here
		maxnum = max740(maxnum, v)
	}

	// take[i] = ar[i] + skip[i-1]
	// skip[i] = max(skip[i-1], take[i-1])
	for i := 1; i <= maxnum; i++ {
		pick[i] = ar[i] + skip[i-1]
		skip[i] = max740(skip[i-1], pick[i-1])
	}
	return max740(pick[maxnum], skip[maxnum])
}

func deleteAndEarnIter2(nums []int) int {
	pick := 0
	skip := 0

	ar := make([]int, 10001)
	maxnum := -1
	for _, v := range nums {
		ar[v] = ar[v] + v // calculating earnings too here
		maxnum = max740(maxnum, v)
	}

	// npick = ar[i] + skip
	// nskip = max(skip, pick)
	// pick,skip = npick, nskip
	for i := 1; i <= maxnum; i++ {
		npick := ar[i] + skip
		nskip := max740(skip, pick)
		pick, skip = npick, nskip
	}
	return max740(pick, skip)
}

func deleteAndEarnRec(nums []int) int {
	sort.Ints(nums)
	freqmap := make(map[int]int)
	newar := make([]int, 0)
	for _, v := range nums {
		if freqmap[v] == 0 {
			newar = append(newar, v)
		}
		freqmap[v]++
	}
	mem := make(map[int]int)
	for k, _ := range freqmap {
		mem[k] = -1
	}
	return delEarnRec(newar, freqmap, 0, mem)
}

func delEarnRec(ar []int, freqmap map[int]int, idx int, mem map[int]int) int {
	if idx >= len(ar) {
		return 0
	}
	val := ar[idx]
	if freqmap[val] > 0 && mem[val] > -1 {
		return mem[val]
	}

	// pick val
	cost1 := 0
	if freqmap[val] > 0 {
		// fmt.Printf("pick %d\n", val)
		cost1 = val * freqmap[val]
		freqnext := freqmap[val+1]
		freqmap[val+1] = 0
		nextidx := idx + 1
		if nextidx < len(ar) && ar[nextidx] == val+1 {
			nextidx++
		}
		cost1 = cost1 + delEarnRec(ar, freqmap, nextidx, mem)
		freqmap[val+1] = freqnext
	}

	// dont pick
	// fmt.Printf("don't pick %d\n", val)
	cost2 := delEarnRec(ar, freqmap, idx+1, mem)
	maxval := max740(cost1, cost2)
	if freqmap[val] > 0 {
		mem[val] = maxval
	}
	// fmt.Printf("mem[%d] = %d\n", val, mem[val])
	return maxval
}

func max740(a, b int) int {
	if a > b {
		return a
	}
	return b
}

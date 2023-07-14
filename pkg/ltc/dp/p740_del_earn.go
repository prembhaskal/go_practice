package dp

import "sort"

func deleteAndEarn(nums []int) int {
	// return deleteAndEarnRec(nums)
	// return deleteAndEarnIter1(nums)
	// return deleteAndEarnIter2(nums)
	// return deleteAndEarnIterLast(nums)
	return deleteAndEarnRec2(nums)
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

// start from last and move towards first
func deleteAndEarnRec2(nums []int) int {
	sort.Ints(nums)
	gain := make(map[int]int)
	ar := make([]int, 0)
	for _, v := range nums {
		if gain[v] == 0 {
			ar = append(ar, v)
		}
		gain[v] = gain[v] + v
	}

	mem := make([]int, len(nums))
	for i := range mem {
		mem[i] = -1
	}

	return deleteAndEarnRecMem2(0, ar, mem, gain)
}

func deleteAndEarnRecMem2(idx int, ar, mem []int, gain map[int]int) int {
	if idx >= len(ar) {
		return 0
	}
	if mem[idx] != -1 {
		return mem[idx]
	}
	curr := ar[idx]
	next := -1
	if idx+1 < len(ar) {
		next = ar[idx+1]
	}

	if curr != next-1 {
		mem[idx] = deleteAndEarnRecMem2(idx+1, ar, mem, gain) + gain[curr]
	} else {
		mem[idx] = max740(deleteAndEarnRecMem2(idx+2, ar, mem, gain)+gain[curr], deleteAndEarnRecMem2(idx+1, ar, mem, gain))
	}
	return mem[idx]
}

func deleteAndEarnIterLast(nums []int) int {
	sort.Ints(nums)
	gain := make(map[int]int)
	ar := make([]int, 0)
	for _, v := range nums {
		if gain[v] == 0 {
			ar = append(ar, v)
		}
		gain[v] = gain[v] + v
	}

	// start from last and move towards first
	// [curr] ...[curr+1, curr+2, ...]

	// if curr != next - 1 , then V[curr] = V[curr+1] + gain[curr] // we can choose both current and next.
	// else , V[curr] = max(V[curr+2], gain[curr]  /*choose current*/ , V[curr+1] /*dont choose current*/  )

	n := len(ar)
	V := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		next := -1
		if i+1 < n {
			next = ar[i+1]
		}
		curr := ar[i]
		if curr != next-1 {
			V[i] = V[i+1] + gain[curr]
		} else {
			V[i] = -1
			if i+2 <= n {
				V[i] = max740(V[i+2]+gain[curr], V[i+1])
			} else {
				V[i] = V[i+1]
			}
		}
	}
	return V[0]
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
	for k := range freqmap {
		mem[k] = -1
	}
	return delEarnRec(newar, freqmap, 0, mem)
}

// this has become too cryptic, check above solutions which are better.
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

package bit

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	numfreq := make(map[int]int)
	uniq := make([]int, 0)

	prev := -100
	for _, v := range nums {
		numfreq[v]++

		if v != prev {
			uniq = append(uniq, v)
			prev = v
		}
	}

	// fmt.Printf("uniq: %v, freqmap: %v\n", uniq, numfreq)

	final := make([][]int, 0)
	curr := make([]int, 0)
	return addSubsetsdup(uniq, numfreq, curr, 0, final)
}

func addSubsetsdup(uniq []int, numfreq map[int]int, curr []int, idx int, final [][]int) [][]int {
	if idx >= len(uniq) {
		final = append(final, curr)
		return final
	}

	// choose freq wise, 1 times, 2 times, 3 times etc...
	num := uniq[idx]
	n := numfreq[num]
	for i := 1; i <= n; i++ {
		dst := make([]int, len(curr))
		copy(dst, curr)
		for f := 0; f < i; f++ {
			dst = append(dst, num)
		}

		final = addSubsetsdup(uniq, numfreq, dst, idx+1, final)
	}

	// don't choose
	return addSubsetsdup(uniq, numfreq, curr, idx+1, final)
}

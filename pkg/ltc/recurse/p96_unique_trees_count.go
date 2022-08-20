package recurse

func numTrees(n int) int {
	// return numTreesRec(1, n)
	return numTreesRecMem1(n)
}

func numTreesRecMem1(n int) int {
	mem := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		mem[i] = make([]int, n+1)
	}

	return numTreesRecMem(1, n, mem)
}

func numTreesRecMem(start, end int, mem [][]int) int {
	if start > end {
		return 0
	}
	if start == end {
		return 1
	}

	if mem[start][end] != 0 {
		return mem[start][end]
	}

	count := 0
	for i := start; i <= end; i++ {
		left := numTreesRecMem(start, i-1, mem)
		right := numTreesRecMem(i+1, end, mem)

		if left == 0 {
			count = count + right
		} else if right == 0 {
			count = count + left
		} else {
			count = count + left*right
		}
	}
	mem[start][end] = count
	return mem[start][end]
}

// generate element with root 1 ... m - 1 , m .. m+1 ... n
//  lefttree root righttree
func numTreesRec(start, end int) int {
	if start > end {
		return 0
	}
	if start == end {
		return 1
	}

	count := 0
	for i := start; i <= end; i++ {
		left := numTreesRec(start, i-1)
		right := numTreesRec(i+1, end)

		if left == 0 {
			count = count + right
		} else if right == 0 {
			count = count + left
		} else {
			count = count + left*right
		}
	}
	return count
}

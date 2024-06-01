package recurse 

import (
	"fmt"
	"math/bits"
)

func minNumberOfSemesters(n int, relations [][]int, k int) int {
	minCnt = 10000

	// create graph
	gph := &Graph{
		adjList: make(map[int][]*Node, 0),
		n:       n,
	}
	nodemap = make(map[int]*Node, 0)
	for i := 1; i <= n; i++ {
		nodemap[i] = &Node{i, 0}
	}
	for _, rel := range relations {
		from := rel[0]
		to := rel[1]

		toNode := nodemap[to]

		gph.adjList[from] = append(gph.adjList[from], toNode)
		toNode.degree++
	}

	n1 := n + 1
	dp1 = make([]int, 1<<n1)
	for i := 0; i < 1<<n1;i++ {
		dp1[i] = -1
	}
	
	minSem := fn2(gph, k, 0, n)
	return minSem
}

type Graph struct {
	adjList map[int][]*Node
	n       int
}

type Node struct {
	id     int
	degree int
}

var dp1 []int
var minCnt = 10000
var nodemap map[int]*Node


// since memoization, always build answer from base (bottom) to up, learned it hard way today.
func fn2(graph *Graph, k, mask, remCourses int) int {
	// fmt.Printf("fn1 call mask: %d\n", mask)
	if dp1[mask] != -1 {
		return dp1[mask]
	}
	// if all courses done, check sem count
	if remCourses == 0 {
		return 0
	}

	// find out available courses
	avail := 0
	tempMask := 0
	for i := 1; i <= graph.n; i++ {
		node := nodemap[i]
			// skip if already picked
			if mask&(1<<i) > 0 {
				continue
			}
			if node.degree == 0 {
				tempMask = tempMask | (1 << i)
				avail++
			}
	}
	if avail == 0 {
		return 0
	}
	
	minSem := 10000

	// try multiple combinations with available nodes with 0 degrees
	allCombs := genCombinations(graph.n, min(k, avail), mask, tempMask)
	for _, comb := range allCombs {
		// try for each comb.
		selected := make([]int, 0)
		for i := 1; i <= graph.n; i++ {
			if comb & (1 <<i) > 0 {
				selected = append(selected, i)
			}
		}

		for _, course := range selected {
			nextnodes := graph.adjList[course]
			for _, next := range nextnodes {
				next.degree--
			}
		}
		// start new sem
		semEnd := fn2(graph, k, mask | comb, remCourses-len(selected)) + 1
		minSem = min(minSem, semEnd)

		// update back the indegree
		for _, course := range selected {
			nextnodes := graph.adjList[course]
			for _, next := range nextnodes {
				next.degree++
			}
		}
	}


	// fmt.Printf("mask: %d, count: %d\n", mask, minSem)

	dp1[mask] = minSem
	// return dp[mask][remCourses]
	return dp1[mask]
}

// all nos from 1 to 1<<n with k bits set
func genCombinations(n, k, mask, andMask int) []int {
	ansmap := make(map[int]bool, 0)
	ans := make([]int, 0)
	last := 1 << (n+1)
	last--
	for i := 1; i <= last; i++ {
		if i & mask > 0 {
			continue
		}
		ni := i & andMask
		if bits.OnesCount32(uint32(ni)) == k {
			ansmap[ni] = true
		}
	}
	for k, _ := range ansmap {
		ans = append(ans, k)
	}
	return ans
}

func genmask(n, mask, curr int) int {
	if curr > n {
		if mask == 16366 {
			fmt.Printf("mask: %.16b\n", mask)
		}
		return 1
	}
	total := 0

	for i := curr; i <=n ;i++ {
		// pick
		total += genmask(n, mask | (1<<i), i + 1)

		// not pick
		total += genmask(n, mask, i + 1)
	}

	return total
}

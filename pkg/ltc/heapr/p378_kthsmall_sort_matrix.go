package heapr

import (
	"container/heap"
	"fmt"
)

func kthSmallest(matrix [][]int, k int) int {
	return kthsmallbyheap(matrix, k)
}

func kthsmallbyheap(mat [][]int, k int) int {
	n := len(mat)
	hcells := cells(make([]*cell, n))
	// load 1st row in heap
	for j := 0; j < n; j++ {
		hcells[j] = &cell{0, j, mat[0][j]}
	}
	heap.Init(&hcells)

	// repeat until kth element found
	for len(hcells) > 0 {
		mincelli := heap.Pop(&hcells)
		mincell := mincelli.(*cell)
		// fmt.Printf("minimum is %s\n", mincell)
		k--
		if k == 0 {
			return mincell.val
		}

		if mincell.i+1 < n {
			// newval := mat[mincell.i+1][mincell.j]
			// fmt.Printf(" inside newval :%d...\n", newval)
			heap.Push(&hcells, &cell{mincell.i + 1, mincell.j, mat[mincell.i+1][mincell.j]})
		}
	}

	return -1
}

type cell struct {
	i   int
	j   int
	val int
}

func (c cell) String() string {
	return fmt.Sprintf("i: %d, j: %d, val: %d\n", c.i, c.j, c.val)
}

type cells []*cell

func (c *cells) Len() int {
	cd := *c
	return len(cd)
}

func (c *cells) Less(i, j int) bool {
	cd := *c
	return cd[i].val < cd[j].val
}

func (c *cells) Swap(i, j int) {
	cd := *c
	cd[i], cd[j] = cd[j], cd[i]
}

func (c *cells) Push(x interface{}) {
	cd := *c
	*c = append(cd, x.(*cell))
}

// Pop return len-1 element
func (c *cells) Pop() interface{} {
	cd := *c
	n := len(cd)
	res := cd[n-1]
	cd[n-1] = nil
	*c = cd[0 : n-1]
	return res
}

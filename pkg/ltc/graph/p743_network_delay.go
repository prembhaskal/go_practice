package graph

import (
	"fmt"
	"math"
)

func networkDelayTime(times [][]int, n int, k int) int {
	// return dijkstraCLRS(times, n, k)
	return dijkstraCLRSHeapInsert(times, n, k)
}

func dijkstraCLRS(times [][]int, n int, k int) int {
	// times[i] = [src, dest, weight]  <- from question
	// n nodes
	// k source

	// init source
	inf := math.MaxInt
	// vtxMap := make(map[int]*vtx743)
	vertices := make([]*vtx743, n+1)
	for i := 1; i <= n; i++ {
		vtx := &vtx743{i, inf}
		// vtxMap[i] = vtx
		vertices[i] = vtx
	}

	adj := make(map[int][]edge743)
	for _, time := range times {
		frm := time[0]
		to := time[1]
		dst := time[2]
		adj[frm] = append(adj[frm], edge743{to: to, dst: dst})
	}

	vertices[k].d = 0

	visited := make(map[int]bool)

	minq := newmheap743(vertices[1:])

	// minq.print()

	// as per CLRS,
	// visited Set
	// add all to Q
	// for Q.is_not_empty
	// u = Q.extractmin
	// for edge (u,v) from adj[u]
	// Relax(u, v, w)
	for minq.size > 0 {
		u := minq.extractmin()
		// fmt.Printf("extract min is %v\n", u)
		visited[u.num] = true
		for _, e := range adj[u.num] {
			if visited[e.to] { // this was not mentioned in CLR
				continue
			}
			// relax(u, v)
			v := vertices[e.to]
			// fmt.Printf("next edge, before: %v\n", v)
			if v.d > u.d+e.dst {
				v.d = u.d + e.dst
				minq.decreasekey(v.num, v.d)
			}
			// fmt.Printf("next edge, after: %v\n", v)
			// minq.print()
		}
	}

	// loop through all to find the max
	max := -1
	for _, vtx := range vertices[1:] {
		if max < vtx.d {
			max = vtx.d
		}
	}

	if max == inf {
		return -1
	}

	return max
}

// On leetcode tests, this is not really that fast, above impl. gives similar time.
func dijkstraCLRSHeapInsert(times [][]int, n int, k int) int {
	// times[i] = [src, dest, weight]  <- from question
	// n nodes
	// k source

	// init source
	inf := math.MaxInt
	// vtxMap := make(map[int]*vtx743)
	vertices := make([]*vtx743, n+1)
	for i := 1; i <= n; i++ {
		vtx := &vtx743{i, inf}
		// vtxMap[i] = vtx
		vertices[i] = vtx
	}

	adj := make(map[int][]edge743)
	for _, time := range times {
		frm := time[0]
		to := time[1]
		dst := time[2]
		adj[frm] = append(adj[frm], edge743{to: to, dst: dst})
	}

	vertices[k].d = 0

	visited := make(map[int]bool)

	// create heap with only source
	minq := newmheap743(nil)
	minq.heapInsert(vertices[k])

	// minq.print()

	// modified CLRS,
	// visited Set
	// add only source to Q
	// for Q.is_not_empty
	// u = Q.extractmin
	// for edge (u,v) from adj[u]
	//   if v not in Q
	//     add to Q, with v.d=INFINITY
	//   Relax(u, v, w)
	for minq.size > 0 {
		u := minq.extractmin()
		// fmt.Printf("extract min is %v\n", u)
		visited[u.num] = true
		for _, e := range adj[u.num] {
			if visited[e.to] { // this was not mentioned in CLR
				continue
			}
			// relax(u, v)
			v := vertices[e.to]
			if !minq.contains(v) {
				minq.heapInsert(v)
			}
			// fmt.Printf("next edge, before: %v\n", v)
			if v.d > u.d+e.dst {
				v.d = u.d + e.dst
				minq.decreasekey(v.num, v.d)
			}
			// fmt.Printf("next edge, after: %v\n", v)
			// minq.print()
		}
	}

	// loop through all to find the max
	max := -1
	for _, vtx := range vertices[1:] {
		if max < vtx.d {
			max = vtx.d
		}
	}

	if max == inf {
		return -1
	}

	return max
}

type edge743 struct {
	to  int
	dst int
}

type vtx743 struct {
	num int
	d   int
}

// stored from 1, last element at ar[size]
type mheap743 struct {
	ar      []*vtx743
	size    int
	itemmap map[int]int // map from edge.v to idx in the array.
}

func newmheap743(items []*vtx743) *mheap743 {
	ar := make([]*vtx743, len(items)+1)
	copy(ar[1:], items)
	q := &mheap743{
		ar:      ar,
		size:    len(items),
		itemmap: make(map[int]int),
	}

	q.heapify()
	return q
}

func (m *mheap743) left(i int) int {
	return 2 * i
}

func (m *mheap743) right(i int) int {
	return 2*i + 1
}

func (m *mheap743) parent(i int) int {
	return i / 2
}

func (m *mheap743) heapify() {
	// init map
	for i := 1; i < len(m.ar); i++ {
		m.itemmap[m.ar[i].num] = i
	}

	// start from last non-leaf node
	for i := m.size / 2; i > 0; i-- {
		m.minHeapify(i)
	}
}

func (m *mheap743) minHeapify(idx int) {
	l := m.left(idx)
	r := m.right(idx)

	min := idx
	if l <= m.size && m.ar[l].d < m.ar[min].d {
		min = l
	}
	if r <= m.size && m.ar[r].d < m.ar[min].d {
		min = r
	}
	if min != idx {
		// swap with min. child
		// m.ar[min], m.ar[idx] = m.ar[idx], m.ar[min]
		m.swap(min, idx)
		// update new places.
		m.updatemap(min)
		m.updatemap(idx)

		// next child
		m.minHeapify(min)
	}
}

func (m *mheap743) updatemap(idx int) {
	m.itemmap[m.ar[idx].num] = idx
}

func (m *mheap743) removemapentry(mapIndex int) {
	m.itemmap[mapIndex] = -1
}

func (m *mheap743) contains(vtx *vtx743) bool {
	if m.itemmap[vtx.num] > 0 {
		return true
	}
	return false
}

func (m *mheap743) extractmin() *vtx743 {
	item := m.ar[1]
	m.ar[1] = m.ar[m.size]
	m.updatemap(1)
	m.removemapentry(item.num)
	// reduce size
	m.ar = m.ar[:m.size]
	m.size--
	m.minHeapify(1)
	return item
}

func (m *mheap743) heapInsert(vtx *vtx743) {
	m.ar = append(m.ar, vtx)
	m.size++
	m.updatemap(m.size)
	origkey := vtx.d
	vtx.d = math.MaxInt
	m.decreasekey(vtx.num, origkey)
}

// we need idx of the key.
func (m *mheap743) decreasekey(vertex int, newdist int) {
	idx := m.itemmap[vertex]
	if m.ar[idx].d < newdist {
		panic("no increase support")
	}

	m.ar[idx].d = newdist

	par := m.parent(idx)
	for idx > 1 && m.ar[idx].d < m.ar[par].d {
		// swap with parent
		m.swap(idx, par)
		m.updatemap(idx)
		m.updatemap(par)
		idx = par
		par = m.parent(idx)
	}
}

func (m *mheap743) swap(i, j int) {
	m.ar[i], m.ar[j] = m.ar[j], m.ar[i]
}

func (m *mheap743) print() {
	fmt.Println("---------- QUEUE ---------")
	for _, vtx := range m.ar {
		fmt.Printf("vtx: %v  ", vtx)
	}
	fmt.Println()
	fmt.Println("---------- QUEUE ---------")
}

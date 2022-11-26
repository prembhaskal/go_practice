package graph

import (
	"fmt"
	"math"
)

func networkDelayTime(times [][]int, n int, k int) int {
	// times[i] = [src, dest, weight]
	// n nodes
	// init source
	inf := math.MaxInt
	vtxMap := make(map[int]*vtx743)
	vertices := make([]*vtx743, 0)
	for i := 1; i <= n; i++ {
		vtx := &vtx743{i, inf}
		vtxMap[i] = vtx
		vertices = append(vertices, vtx)
		// fmt.Printf("add vertex: %v\n", vtx)
	}

	// wgtmap := make(map[edge743]int)
	adj := make(map[int][]edge743)
	for _, time := range times {
		frm := time[0]
		to := time[1]
		dst := time[2]
		adj[frm] = append(adj[frm], edge743{to: to, dst: dst})
		// wgtmap[edge743{frm, to}] = dst
	}

	vtxMap[k].d = 0

	visited := make(map[int]bool)

	minq := newmheap743(vertices)

	// minq.print()

	for minq.size > 0 {
		u := minq.extractmin()
		// fmt.Printf("extract min is %v\n", u)
		visited[u.num] = true
		for _, e := range adj[u.num] {
			if visited[e.to] {
				continue
			}
			// relax(u, v)
			v := vtxMap[e.to]
			// fmt.Printf("next edge, before: %v\n", v)
			// if v.d > u.d+wgtmap[edge743{u.num, v.num}] {
			// 	v.d = u.d + wgtmap[edge743{u.num, v.num}]
			// 	minq.decreasekey(v.num, v.d)
			// }
			if v.d > u.d + e.dst {
				v.d = u.d + e.dst
				minq.decreasekey(v.num, v.d)
			}
			// fmt.Printf("next edge, after: %v\n", v)
			// minq.print()
		}
	}

	// visited Set
	// add all to Q
	// for Q.is_not_empty
	// u = Q.extractmin
	// for edge (u,v) from adj[u]
	// Relax(u, v, w)

	// loop through all to find the max
	max := -1
	for _, vtx := range vertices {
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
	from int
	to   int
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

func (m *mheap743) extractmin() *vtx743 {
	item := m.ar[1]
	m.ar[1] = m.ar[m.size]
	m.updatemap(1)
	// reduce size
	m.ar = m.ar[:m.size]
	m.size--
	m.minHeapify(1)
	return item
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

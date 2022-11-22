package graph

import "sort"

// TODO - implement prim's algorithm with vertices in heap instead of edges.
func minCostConnectPoints(points [][]int) int {
	graph := make(map[int][]edge1584)
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1 := points[i]
			p2 := points[j]
			dist := mhtdist(p1, p2)
			graph[i] = append(graph[i], edge1584{i, j, dist})
			graph[j] = append(graph[j], edge1584{j, i, dist})
		}
	}

	// prims algo
	visited := make(map[int]bool)
	// choose a vertex, add all its edges to min-heap
	heap := newminheap1584()
	visited[0] = true
	for _, edge := range graph[0] {
		heap.add(edge)
	}

	mst := 0
	mstcnt := n - 1
	for mstcnt > 0 {
		minedge := heap.poll()
		to := minedge.to
		if visited[to] {
			continue
		}
		visited[to] = true
		mst += minedge.dist
		mstcnt--

		for _, edge := range graph[to] {
			if !visited[edge.to] {
				heap.add(edge)
			}
		}
	}

	return mst
}

type minheap1584 struct {
	ar  []edge1584
	siz int
}

func newminheap1584() *minheap1584 {
	return &minheap1584{
		ar: make([]edge1584, 1), // 1 based index.
	}
}

func (h *minheap1584) left(i int) int {
	return 2 * i
}

func (h *minheap1584) right(i int) int {
	return 2*i + 1
}

func (h *minheap1584) parent(i int) int {
	return i / 2
}

func (h *minheap1584) heapify(i int) {
	l := h.left(i)
	r := h.right(i)
	min := i
	if l <= h.siz && h.ar[l].dist < h.ar[i].dist {
		min = l
	}
	if r <= h.siz && h.ar[r].dist < h.ar[min].dist {
		min = r
	}
	if min != i {
		h.ar[min], h.ar[i] = h.ar[i], h.ar[min]
		h.heapify(min)
	}
}

func (h *minheap1584) poll() edge1584 {
	if h.siz == 0 {
		panic("empty heap")
	}
	item := h.ar[1]
	h.ar[1] = h.ar[h.siz]
	h.siz--
	h.ar = h.ar[:h.siz+1]
	h.heapify(1)
	return item
}

// 0 1 2 3  size = 3
// 0 1 2  size = 2

func (h *minheap1584) add(item edge1584) {
	h.ar = append(h.ar, item)
	h.siz++
	i := h.siz
	for {
		par := h.parent(i)
		if h.ar[i].dist < h.ar[par].dist {
			h.ar[i], h.ar[par] = h.ar[par], h.ar[i]
			i = par
		} else {
			break
		}
	}
}

func (h *minheap1584) length() int {
	return h.siz
}

func minCostConnectPoints1(points [][]int) int {
	edges := make([]edge1584, 0)
	// find all edges
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1 := points[i]
			p2 := points[j]
			dist := mhtdist(p1, p2)
			// fmt.Printf("p1: %v, p2: %v, dist: %d\n", p1, p2, dist)
			edges = append(edges, edge1584{i, j, dist})
			edges = append(edges, edge1584{j, i, dist})
		}
	}

	// run kruskal's algo

	// sort edges
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	// fmt.Printf("sorted edges: %v\n", edges)

	ufind := newunionfind(len(edges))

	mst := 0

	for i := 0; i < len(edges); i++ {
		rt1 := ufind.find(edges[i].from)
		rt2 := ufind.find(edges[i].to)

		if rt1 != rt2 {
			ufind.union(edges[i].from, edges[i].to)
			mst = mst + edges[i].dist
		}
	}

	return mst
}

type ufind struct {
	root []int
	rank []int
}

func newunionfind(n int) *ufind {
	root := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
	}
	return &ufind{
		root: root,
		rank: make([]int, n),
	}
}

func (u *ufind) find(x int) int {
	if x == u.root[x] {
		return x
	}
	rt := u.find(u.root[x])
	u.root[x] = rt
	return rt
}

func (u *ufind) union(x, y int) {
	rx := u.find(x)
	ry := u.find(y)
	if rx != ry {
		if u.rank[rx] > u.rank[ry] {
			u.root[ry] = rx
		} else if u.rank[ry] > u.rank[rx] {
			u.root[rx] = ry
		} else {
			u.root[rx] = ry
			u.rank[ry]++
		}
	}
}

type edge1584 struct {
	from int
	to   int
	dist int
}

func mhtdist(p1, p2 []int) int {
	return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

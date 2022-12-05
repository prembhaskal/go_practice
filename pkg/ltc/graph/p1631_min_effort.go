package graph

import "fmt"

// TODO - add the necessary stuff to compile below stuff
// TODO - solve using binary search approach too.
// TODO - solve using BFS too.
func minimumEffortPath(heights [][]int) int {
	// You are given heights, a 2D array of size rows x columns,
	// where heights[row][col] represents the height of cell (row, col).
	// return -1

	rows := len(heights)
	cols := len(heights[0])

	// fmt.Printf("rows: %d, cols: %d\n", rows, cols)

	minq := newminheap1631()

	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	inf := 10000000
	vertices := make([][]*vtx1631, 0) // m x n size
	for i := 0; i < rows; i++ {
		vertices = append(vertices, make([]*vtx1631, cols))
		for j := 0; j < cols; j++ {
			vertices[i][j] = &vtx1631{
				row: i,
				col: j,
				d:   inf,
			}
		}
	}

	visited := make(map[pair1631]bool)

	vtxzero := vertices[0][0]
	vtxzero.d = 0
	minq.insert(vtxzero)

	for minq.size > 0 {
		u := minq.extractMin()
		visited[pair1631{u.row, u.col}] = true

		currht := heights[u.row][u.col]

		for _, dir := range dirs {
			nr := u.row + dir[0]
			nc := u.col + dir[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				continue
			}

			if visited[pair1631{nr, nc}] {
				continue
			}

			v := vertices[nr][nc]

			if !minq.contains(pair1631{nr, nc}) {
				minq.insert(v)
			}

			// if abs1631(heights[nr][nc] - currht) > u.d

			// if v.d > u.d+abs1631(heights[nr][nc]-currht) {
			// 	v.d = u.d + abs1631(heights[nr][nc]-currht)
			// 	minq.decreasekeypair(pair1631{nr, nc}, v.d)
			// }

			if v.d > max1631(abs1631(heights[nr][nc]-currht), u.d) {
				newht := max1631(abs1631(heights[nr][nc]-currht), u.d)
				// fmt.Printf("in decrease key, row: %d, col: %d, exist: %d, new: %d\n", nr, nc, v.d, newht)
				v.d = newht
				minq.decreasekeypair(pair1631{nr, nc}, v.d)
			}
		}

	}

	return vertices[rows-1][cols-1].d
}

func max1631(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs1631(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type vtx1631 struct {
	row int
	col int
	d   int
}

type pair1631 struct {
	row int
	col int
}

type minheap1631 struct {
	ar     []*vtx1631 // we store from 1 onwards, last element index == size.
	size   int
	idxmap map[pair1631]int
}

func newminheap1631() *minheap1631 {
	return &minheap1631{
		ar:     make([]*vtx1631, 1),
		size:   0,
		idxmap: make(map[pair1631]int),
	}
}

func (q *minheap1631) parent(i int) int {
	return i / 2
}

func (q *minheap1631) left(i int) int {
	return 2 * i
}

func (q *minheap1631) right(i int) int {
	return 2*i + 1
}

func (q *minheap1631) insert(item *vtx1631) {
	q.ar = append(q.ar, item)
	q.size++

	q.updatemap(q.size)
	key := item.d
	item.d = 10000000
	q.decreaseKey(q.size, key)
}

func (q *minheap1631) decreasekeypair(cell pair1631, key int) {
	var idx int
	var ok bool
	if idx, ok = q.idxmap[cell]; !ok {
		panic(fmt.Sprintf("key not present in map: %v", cell))
	}
	q.decreaseKey(idx, key)
}

func (q *minheap1631) decreaseKey(idx int, key int) {
	item := q.ar[idx]
	if item.d < key {
		fmt.Printf("decrease panic item row: %d, col: %d, exist key: %d, new key: %d\n", item.row, item.col, item.d, key)
		fmt.Printf("index map content: %v\n", q.idxmap)
		fmt.Printf("array content: %v", q.ar[1:])
		panic("only decrease supported")
	}
	item.d = key
	par := q.parent(idx)
	for idx > 1 && q.ar[idx].d < q.ar[par].d {
		q.swap(idx, par)
		q.updatemap(idx)
		q.updatemap(par)
		idx = par
		par = q.parent(idx)
	}
}

func (q *minheap1631) swap(i, j int) {
	q.ar[i], q.ar[j] = q.ar[j], q.ar[i]
}

func (q *minheap1631) extractMin() *vtx1631 {
	if q.size <= 0 {
		panic("empty queue")
	}
	item := q.ar[1]
	q.ar[1] = q.ar[q.size] // put last to first.

	// update map
	q.updatemap(1)
	delete(q.idxmap, pair1631{item.row, item.col})

	q.ar = q.ar[:q.size]

	q.size--

	q.heapify(1)
	return item
}

func (q *minheap1631) heapify(idx int) {
	left := q.left(idx)
	right := q.right(idx)
	min := idx
	if left <= q.size && q.ar[left].d < q.ar[min].d {
		min = left
	}
	if right <= q.size && q.ar[right].d < q.ar[min].d {
		min = right
	}
	if min != idx {
		q.swap(min, idx)
		q.updatemap(min)
		q.updatemap(idx)
		q.heapify(min)
	}
}

func (q *minheap1631) updatemap(idx int) {
	q.idxmap[pair1631{q.ar[idx].row, q.ar[idx].col}] = idx
}

func (q *minheap1631) contains(cell pair1631) bool {
	if _, ok := q.idxmap[cell]; !ok {
		return false
	}
	return true
}
func (q *minheap1631) f() {}

package graph

// TODO - add the necessary stuff to compile below stuff
// TODO - solve using binary search approach too.
// TODO - solve using BFS too.
func minimumEffortPath(heights [][]int) int {
	return -1
	rows := len(heights)
	cols := len(heights[0])

	minq := newminheap1631();

	dirs := [][]int{{-1,0}, {1,0}, {0,1}, {0,-1}}

	vertices := make([][]*vtx1631, 0) // m x n size

	for minq.size > 0 {
		u := minq.extractmin()
		visited[pair1631{u.row, u.col}] = true

		currht := heights[u.row][u.col]

		for _, dir := range dirs {
			nr := u.row + dir[0]
			nc := u.col + dir[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				continue
			}

			if visited[{nr, nc}] {
				continue
			}

			if !minq.contains({nr, nc}) {
				minq.insert(v)
			}

			v := vertices[nr][nc]
			if v.d > u.d + abs(heights[nr][nc] - currht) {
				v.d = u.d + abs(heights[nr][nc] - currht)
				minq.decreasekey({nr, nc}, v.d)
			}

		}

	}

	return vertices[rows-1][cols-1].d
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
		ar:   make([]*vtx1631, 1),
		size: 0,
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

func (q *minheap1631) decreaseKey(idx int, key int) {
	item := q.ar[idx]
	if item.d < key {
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
	item := q.ar[0]
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
		q.heapify(min)
	}
}

func (q *minheap1631) updatemap(idx int) {
	q.idxmap[pair1631{q.ar[idx].row, q.ar[idx].col}] = idx
}

func (q *minheap1631) contains(row, col int) bool {
	if _, ok := q.idxmap[pair1631{row, col}]; !ok {
		return false
	}
	return true
}
func (q *minheap1631) f() {}

package graph

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)

	cells := make([][]*cell, 0)

	for i := 0; i < n; i++ {
		row := make([]*cell, 0)
		for j := 0; j < n; j++ {
			row = append(row, &cell{i, j, grid[i][j], -1, false})
		}

		cells = append(cells, row)
	}

	q := newqueue1091()

	if cells[0][0].val == 0 {
		cells[0][0].lev = 1
		q.add(cells[0][0])
	}

	for q.size() > 0 {
		curr := q.poll()
		curr.seen = true

		// fmt.Printf("curr cell: %v\n", curr)

		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				if x == 0 && y == 0 {
					continue
				}
				ni := x + curr.i
				nj := y + curr.j

				// fmt.Printf("ni: %d, nj: %d\n", ni, nj)

				if ni >= 0 && ni < n && nj >= 0 && nj < n {
					next := cells[ni][nj]
					if !next.seen && next.val == 0 {
						next.seen = true
						next.lev = curr.lev + 1
						q.add(next)
						// fmt.Printf("next cell: %v\n", next)
					}
				}
			}
		}
	}

	return cells[n-1][n-1].lev
}

type cell struct {
	i    int
	j    int
	val  int
	lev  int
	seen bool
}

type queue1091 struct {
	ar []*cell
}

func newqueue1091() *queue1091 {
	return &queue1091{
		ar: make([]*cell, 0),
	}
}

func (q *queue1091) size() int {
	return len(q.ar)
}

func (q *queue1091) poll() *cell {
	if q.size() == 0 {
		panic("empty queue")
	}
	item := q.ar[0]
	q.ar = q.ar[1:]
	return item
}

func (q *queue1091) add(item *cell) {
	q.ar = append(q.ar, item)
}

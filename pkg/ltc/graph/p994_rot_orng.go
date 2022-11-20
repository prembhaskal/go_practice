package graph

import "fmt"

func orangesRotting(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	q := newqueue994()

	cells := make([][]*cell994, 0)

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < n; i++ {
		row := make([]*cell994, 0)
		for j := 0; j < m; j++ {
			cell := &cell994{i, j, grid[i][j], 0}
			if grid[i][j] == 2 {
				q.add(cell)
			}
			row = append(row, cell)
		}
		cells = append(cells, row)
	}

	for q.size() > 0 {
		curr := q.poll()
		// fmt.Printf("curr: %v\n", curr)
		for _, xx := range dirs {
			ni := curr.i + xx[0]
			nj := curr.j + xx[1]
			if ni >= 0 && ni < n && nj >= 0 && nj < m {
				next := cells[ni][nj]
				if next.val == 1 { // don't need seen flag, since all rotten are already added above.
					next.timetorot = curr.timetorot + 1
					next.val = 2
					// fmt.Printf("next: %v\n", next)
					q.add(next)
				}
			}

		}
	}

	// printcells994(cells)

	// find if all are done, also find oldest time to rot
	latestTime := -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if cells[i][j].val == 1 {
				return -1
			}
			if latestTime < cells[i][j].timetorot {
				latestTime = cells[i][j].timetorot
			}
		}
	}

	return latestTime
}

func printcells994(cells [][]*cell994) {
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[i]); j++ {
			fmt.Printf("[%d, %d : %d]", i, j, cells[i][j].val)
		}
	}
}

type cell994 struct {
	i         int
	j         int
	val       int
	timetorot int
}

type queue994 struct {
	ar []*cell994
}

func newqueue994() *queue994 {
	return &queue994{
		ar: make([]*cell994, 0),
	}
}

func (q *queue994) add(item *cell994) {
	q.ar = append(q.ar, item)
}

func (q *queue994) poll() *cell994 {
	if q.size() == 0 {
		panic("empty queue")
	}
	item := q.ar[0]
	q.ar = q.ar[1:]
	return item
}

func (q *queue994) size() int {
	return len(q.ar)
}

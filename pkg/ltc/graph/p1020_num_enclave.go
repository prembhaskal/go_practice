package graph

func numEnclaves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		bfs(i, 0, grid)
		bfs(i, n-1, grid)
	}
	for j := 0; j < n; j++ {
		bfs(0, j, grid)
		bfs(m-1, j, grid)
	}

	total := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			total += grid[i][j]
		}
	}

	return total
}

func bfs(i, j int, grid [][]int) {
	qu := newmyqueue1020()
	qu.enqueue(mycell1020{i, j})

	for qu.size() > 0 {
		cell := qu.dequeue()
		if cell.i < 0 || cell.j < 0 || cell.i >= len(grid) || cell.j >= len(grid[0]) {
			continue
		}
		if grid[cell.i][cell.j] == 0 {
			continue
		}
		grid[cell.i][cell.j] = 0
		neigh := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
		for _, v := range neigh {
			qu.enqueue(mycell1020{cell.i + v[0], cell.j + v[1]})
		}
	}
}

type myqueue1020 struct {
	ar []mycell1020
}

func newmyqueue1020() *myqueue1020 {
	return &myqueue1020{
		ar: make([]mycell1020, 0),
	}
}

func (m *myqueue1020) enqueue(cell mycell1020) {
	m.ar = append(m.ar, cell)
}

func (m *myqueue1020) dequeue() mycell1020 {
	if m.size() == 0 {
		panic("dequeue on empty queue")
	}
	cell := m.ar[0]
	m.ar = m.ar[1:]
	return cell
}

func (m *myqueue1020) size() int {
	return len(m.ar)
}

type mycell1020 struct {
	i int
	j int
}

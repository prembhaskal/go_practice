package queue_stack

func numIslands(grid [][]byte) int {
	return doDFSWithStack(grid)
}

func doDFSWithStack(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	visited := make([][]bool, len(grid))
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	total := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cell := islandcell{i, j}
			if string(grid[i][j]) == "1" && !visited[i][j] {
				stk := newislandstack()
				stk.push(cell)
				doDFSStack(visited, grid, stk)
				total++
			}
		}
	}

	return total
}

func doDFSStack(visited [][]bool, grid [][]byte, stk *islandstack) {
	for !stk.isempty() {
		cell := stk.pop()
		if visited[cell.i][cell.j] {
			continue
		}
		visited[cell.i][cell.j] = true
		addToStack(stk, grid, cell.i+1, cell.j)
		addToStack(stk, grid, cell.i-1, cell.j)
		addToStack(stk, grid, cell.i, cell.j+1)
		addToStack(stk, grid, cell.i, cell.j-1)
	}
}

func addToStack(stk *islandstack, grid [][]byte, i, j int) {
	if i >= len(grid) || i < 0 {
		return
	}
	if j >= len(grid[0]) || j < 0 {
		return
	}

	if string(grid[i][j]) == "0" {
		return
	}

	stk.push(islandcell{i, j})
}

type islandcell struct {
	i int
	j int
}

type islandstack struct {
	arr []islandcell
}

func newislandstack() *islandstack {
	return &islandstack{
		arr: make([]islandcell, 0),
	}
}

func (s *islandstack) push(cell islandcell) {
	s.arr = append(s.arr, cell)
}

func (s *islandstack) pop() islandcell {
	var icell islandcell
	if s.isempty() {
		return icell
	}

	icell = s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return icell
}

func (s *islandstack) isempty() bool {
	return len(s.arr) == 0
}

func doBFS(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])

	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	total := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if visited[i][j] {
				continue
			}
			if string(grid[i][j]) == "0" {
				continue
			}
			searchByBFS(grid, visited, i, j)
			total++
		}
	}

	return total

}

func searchByBFS(grid [][]byte, visited [][]bool, i, j int) {
	nq := newqueue()
	nq.add(newnode(i, j))

	for !nq.isempty() {
		nd, _ := nq.poll()
		if visited[nd.i][nd.j] {
			continue
		}
		if string(grid[nd.i][nd.j]) == "0" {
			continue
		}
		visited[nd.i][nd.j] = true

		addToQueue(grid, nd.i+1, nd.j, nq)
		addToQueue(grid, nd.i-1, nd.j, nq)
		addToQueue(grid, nd.i, nd.j+1, nq)
		addToQueue(grid, nd.i, nd.j-1, nq)
	}

}

func addToQueue(grid [][]byte, i, j int, nq *nodequeue) {
	if i < 0 || i >= len(grid) {
		return
	}
	if j < 0 || j >= len(grid[0]) {
		return
	}
	if string(grid[i][j]) == "0" {
		return
	}
	nq.add(newnode(i, j))
	return
}

type node struct {
	i int
	j int
}

func newnode(i, j int) node {
	return node{i, j}
}

type nodequeue struct {
	arr []node
}

func newqueue() *nodequeue {
	return &nodequeue{
		arr: make([]node, 0),
	}
}

func (q *nodequeue) add(item node) {
	q.arr = append(q.arr, item)
}

func (q *nodequeue) poll() (*node, bool) {
	if q.isempty() {
		return nil, false
	}
	item := q.arr[0]
	q.arr = q.arr[1:]
	return &item, true
}

func (q *nodequeue) isempty() bool {
	return len(q.arr) == 0
}

func doDFS(grid [][]byte) int {

	total := 0
	seen := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		seen[i] = make([]bool, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if string(grid[i][j]) == "1" && !seen[i][j] {
				// fmt.Printf("checking node %d:%d\n", i , j)
				// seen[i][j] = true
				visitConnected(i, j, grid, seen)
				total++
			}
		}
	}

	return total

}

func visitConnected(i, j int, grid [][]byte, seen [][]bool) {
	// fmt.Printf("visiting %d:%d\n", i , j)
	// fmt.Printf("len(Grid): %d\n", len(grid))
	// fmt.Printf("len(Grid[0]): %d\n", len(grid[0]))
	if i < 0 || i >= len(grid) {
		return
	}

	if j < 0 || j >= len(grid[0]) {
		return
	}

	if seen[i][j] {
		return
	}

	seen[i][j] = true

	if string(grid[i][j]) == "0" {
		return
	}

	visitConnected(i-1, j, grid, seen)
	visitConnected(i+1, j, grid, seen)
	visitConnected(i, j-1, grid, seen)
	visitConnected(i, j+1, grid, seen)

}

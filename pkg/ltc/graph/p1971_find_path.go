package graph

func validPath(n int, edges [][]int, source int, destination int) bool {
	// adj . list
	graph := make([][]int, n)

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]

		lst := graph[u]
		lst = append(lst, v)
		graph[u] = lst

		rlst := graph[v]
		rlst = append(rlst, u)
		graph[v] = rlst
	}

	visited := make([]bool, n)

	queue := newp1971queue()
	queue.add(source)

	for queue.size() > 0 {
		vert := queue.poll()
		if visited[vert] {
			continue
		}

		if vert == destination {
			return true
		}

		visited[vert] = true

		for _, neigh := range graph[vert] {
			queue.add(neigh)
		}
	}

	return false
}

type p1971queue struct {
	ar []int
}

func newp1971queue() *p1971queue {
	return &p1971queue{
		ar: make([]int, 0),
	}
}

func (q *p1971queue) add(i int) {
	q.ar = append(q.ar, i)
}

func (q *p1971queue) poll() int {
	if q.size() == 0 {
		panic("poll on empty queue")
	}

	val := q.ar[0]
	q.ar = q.ar[1:]
	return val
}

func (q *p1971queue) size() int {
	return len(q.ar)
}

func validPath1(n int, edges [][]int, source int, destination int) bool {
	// adj . list
	graph := make([][]int, n)

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]

		lst := graph[u]
		lst = append(lst, v)
		graph[u] = lst

		rlst := graph[v]
		rlst = append(rlst, u)
		graph[v] = rlst
	}

	visited := make([]bool, n)
	return dfs(source, destination, graph, visited)
}

func dfs(curr, dst int, graph [][]int, visited []bool) bool {
	if curr == dst {
		return true
	}
	if visited[curr] {
		return false
	}

	visited[curr] = true

	neighs := graph[curr]
	for _, neigh := range neighs {
		found := dfs(neigh, dst, graph, visited)
		if found {
			return true
		}
	}

	return false
}

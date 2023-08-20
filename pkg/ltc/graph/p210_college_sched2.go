package graph

// Topological sort
func findOrder(numCourses int, prerequisites [][]int) []int {
	// return tsortKahnAlgo(numCourses, prerequisites)
	return tsortDFS(numCourses, prerequisites)
}

func tsortDFS(n int, prereq [][]int) []int {
	adjlist := make([][]int, n)
	for _, edge := range prereq {
		from := edge[0]
		to := edge[1]
		adjlist[from] = append(adjlist[from], to)
	}

	tsort := make([]int, n)

	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs210(i, adjlist, visited, tsort)
		}
	}

	if len(tsort) != n {
		return nil
	}
	return rev210(tsort)
}

func rev210(x []int) []int {
	i := 0
	j := len(x) - 1
	for i < j {
		x[i], x[j] = x[j], x[i]
	}
	return x
}

func dfs210(x int, adjlist [][]int, visited []bool, tsort []int) {
	visited[x] = true
	for _, next := range adjlist[x] {
		if !visited[next] {
			dfs210(next, adjlist, visited, tsort)
		}
	}

	tsort = append(tsort, x)
}

func tsortKahnAlgo(numCourses int, prerequisites [][]int) []int {
	adjList := make([][]int, numCourses)
	indeg := make([]int, numCourses)

	for _, edge := range prerequisites {
		from := edge[1]
		to := edge[0]
		adjList[from] = append(adjList[from], to)
		indeg[to]++
	}

	q := newqueue()

	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q.add(i)
		}
	}

	visited := make([]bool, numCourses)
	tsort := make([]int, 0)

	for q.size() > 0 {
		item := q.poll()
		tsort = append(tsort, item)
		visited[item] = true
		for _, next := range adjList[item] {
			if !visited[next] {
				indeg[next]--
				if indeg[next] == 0 {
					q.add(next)
				}
			}
		}
	}

	if len(tsort) != numCourses {
		return nil
	}

	return tsort
}

type queue struct {
	ar []int
}

func newqueue() *queue {
	return &queue{
		ar: make([]int, 0),
	}
}

func (q *queue) size() int {
	return len(q.ar)
}

func (q *queue) poll() int {
	if q.size() == 0 {
		panic("empty queue")
	}
	item := q.ar[0]
	q.ar = q.ar[1:]
	return item
}

func (q *queue) add(item int) {
	q.ar = append(q.ar, item)
}

func (q *queue) f() {}

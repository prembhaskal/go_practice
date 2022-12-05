package graph

func findOrder(numCourses int, prerequisites [][]int) []int {
	return tsortKhanAlgo(numCourses, prerequisites)
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

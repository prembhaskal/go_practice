package graph

func allPathsSourceTarget(graph [][]int) [][]int {
	// graph is already adj. list

	// source 0
	// dest   n-1
	// if we reach n-1, that means all previous levels are in path, we can maintain that in a stack or array.  that will path array.
	//

	dest := len(graph) - 1
	// pseudo code

	// get curr node from queue
	//   add
	// for curr node, list of neighbours
	// for each neighbours,
	// if neighbour == dest
	//   add currpath to result

	result := make([][]int, 0)

	// return result.
	queue := newp797q()
	source := path797{node: 0, prevpath: []int{0}}
	queue.add(source)

	for queue.size() > 0 {
		currpath := queue.poll()
		currnode := currpath.node
		prevpath := currpath.prevpath

		if currnode == dest {
			result = append(result, prevpath)
		}

		nextnodes := graph[currnode]
		for _, next := range nextnodes {
			// copy of currpath
			copypath := make([]int, len(prevpath))
			copy(copypath, prevpath)
			copypath = append(copypath, next)
			nextpath := path797{node: next, prevpath: copypath}

			queue.add(nextpath)
		}
	}

	return result
}

type path797 struct {
	node     int
	prevpath []int
}

type p797q struct {
	ar []path797
}

func newp797q() *p797q {
	return &p797q{
		ar: make([]path797, 0),
	}
}

func (q *p797q) add(item path797) {
	q.ar = append(q.ar, item)
}

func (q *p797q) poll() path797 {
	if q.size() == 0 {
		panic("empty poll")
	}
	first := q.ar[0]
	q.ar = q.ar[1:]
	return first
}

func (q *p797q) size() int {
	return len(q.ar)
}

func allPathsSourceTarget1(graph [][]int) [][]int {
	n := len(graph)
	allpaths := make([][]int, 0)

	currpath := []int{0}
	return dfsGetPath1(graph, 0, n-1, currpath, allpaths)
}

// no visited needed since graph is DAG
func dfsGetPath(graph [][]int, curr, dest int, currpath []int, allpaths [][]int) [][]int {
	if curr == dest {
		newpath := make([]int, len(currpath))
		copy(newpath, currpath)
		allpaths = append(allpaths, newpath)
		return allpaths
	}

	neighs := graph[curr]
	for _, neigh := range neighs {
		newpath := make([]int, len(currpath))
		copy(newpath, currpath)
		newpath = append(newpath, neigh)

		allpaths = dfsGetPath(graph, neigh, dest, newpath, allpaths)
	}

	return allpaths
}

func dfsGetPath1(graph [][]int, curr, dest int, currpath []int, allpaths [][]int) [][]int {
	if curr == dest {
		newpath := make([]int, len(currpath))
		copy(newpath, currpath)
		allpaths = append(allpaths, newpath)
		return allpaths
	}

	neighs := graph[curr]
	for _, neigh := range neighs {
		currpath = append(currpath, neigh)
		allpaths = dfsGetPath1(graph, neigh, dest, currpath, allpaths)
		currpath = currpath[:len(currpath)-1]
	}

	return allpaths
}

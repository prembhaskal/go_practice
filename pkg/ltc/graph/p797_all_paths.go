package graph

func allPathsSourceTarget(graph [][]int) [][]int {
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

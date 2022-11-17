package graph

func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
	// make adj list
	graph := make(map[int][]int)

	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		graph[from] = append(graph[from], to)
	}

	visited := make([]int, n)
	// visited[source] = 1

	return pathexistdfs(graph, visited, source, destination)
}

// GRAPH coloring - 0 (not processed), 1 (processing), 2 (processed)
func pathexistdfs(graph map[int][]int, visited []int, src, dest int) bool {
	if visited[src] != 0 {
		if visited[src] == 1 {
			return false // this is cycle
		}
		return true // already processed earlier, it will dest. from there.
	}

	if src == dest {
		if len(graph[dest]) > 0 {
			return false
		}
		return true
	}

	visited[src] = 1

	nextnodes := graph[src]
	if len(nextnodes) == 0 {
		return false
	}
	for _, next := range nextnodes {
		if visited[next] == 1 {
			return false
		}

		exist := pathexistdfs(graph, visited, next, dest)
		if !exist {
			return false
		}
	}

	visited[src] = 2 // node processed

	return true
}

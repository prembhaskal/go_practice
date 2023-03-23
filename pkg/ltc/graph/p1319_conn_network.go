package graph

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	// create adjacency list
	graph := make([][]int, n)
	for _, edge := range connections {
		from := edge[0]
		to := edge[1]
		// add from->to
		graph[from] = append(graph[from], to)
		// add to->from
		graph[to] = append(graph[to], from)
	}

	// do dfs
	visited := make([]bool, n)
	scccount := 0
	for node := 0; node < n; node++ {
		if !visited[node] {
			scccount++
			visited[node] = true
			dfs1319(visited, node, graph)
		}
	}

	// needed edges = scc-1
	return scccount - 1
}

func dfs1319(visited []bool, node int, graph [][]int) {
	neighs := graph[node]
	for _, neigh := range neighs {
		if !visited[neigh] {
			// mark connection as visited
			// markPair(node, neigh, connpairs)
			visited[neigh] = true
			dfs1319(visited, neigh, graph)
		}
	}
}

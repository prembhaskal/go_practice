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

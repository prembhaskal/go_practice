package graph

// Approach - use dfs with 3 colors, 0-white,1-grey,2-black to find a cycle
// this is cycle detection in a directed graph.
func canFinish(numCourses int, prerequisites [][]int) bool {
	n := numCourses
	pre := prerequisites

	// make graph
	graph := make([][]int, n)
	for i := 0; i < len(pre); i++ {
		from := pre[i][0]
		to := pre[i][1]
		nexts := graph[from]
		nexts = append(nexts, to)
		graph[from] = nexts
	}

	hasCycle = false
	visited := make([]int, n)
	for i := 0; i < n; i++ {
		dfsColor(i, visited, graph)
	}
	return !hasCycle
}

var hasCycle bool

func dfsColor(curr int, visited []int, graph [][]int) {
	if visited[curr] == 2 {
		return
	}
	if visited[curr] == 1 {
		// fmt.Printf("cycle --> node: %d already marked grey\n", curr)
		hasCycle = true
		return
	}

	visited[curr] = 1 // mark grey
	for _, next := range graph[curr] {
		dfsColor(next, visited, graph)
	}
	visited[curr] = 2 // mark black
}

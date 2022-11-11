package graph

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// create graph from equations
	graph := make(map[string][]edge)
	i := 0
	for _, eq := range equations {
		A := eq[0]
		B := eq[1]
		A_B := values[i]
		B_A := 1.0 / A_B
		edges := graph[A]
		edges = append(edges, edge{A, B, A_B})
		graph[A] = edges

		redges := graph[B]
		redges = append(redges, edge{B, A, B_A})
		graph[B] = redges
		// fmt.Printf("graph at stageL %d is %v\n", i, graph)
		i++
	}

	// fmt.Printf("graph created is %v\n", graph)

	ans := make([]float64, 0)
	for _, query := range queries {
		divans := calcDiv(graph, query[0], query[1])
		ans = append(ans, divans)
	}

	return ans
}

func calcDiv(graph map[string][]edge, from, to string) float64 {
	// check if graph has vertice
	_, okfrom := graph[from]
	_, okto := graph[to]

	if !okfrom {
		return -1.0
	}

	if !okto {
		return -1.0
	}

	if from == to {
		return 1.0
	}
	visited := make(map[string]bool)
	visited[from] = true
	return dfsDiv(graph, from, to, visited)
}

func dfsDiv(graph map[string][]edge, from, to string, visited map[string]bool) float64 {
	// find neighbours
	neighs := graph[from]
	for _, neigh := range neighs {
		if visited[neigh.b] {
			continue
		}

		visited[neigh.b] = true

		if neigh.b == to {
			return neigh.weight
		}

		weight := dfsDiv(graph, neigh.b, to, visited)
		if weight > 0 {
			return weight * neigh.weight
		}
	}

	return -1.0
}

// type graph struct {
//     adjlist map[string][]edge
// }

type edge struct {
	a      string
	b      string
	weight float64
}

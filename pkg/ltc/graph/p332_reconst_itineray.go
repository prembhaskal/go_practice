package graph

import (
	"fmt"
	"sort"
)

// TODO - do using https://www.geeksforgeeks.org/hierholzers-algorithm-directed-graph/ algorithm.
// Current approach - backtracking + greedy ( discard bad solutions early)
func findItinerary(tickets [][]string) []string {
	//     JFK -> [ATL|cross,  SFO|cross]

	graph := make(map[string][]string)
	visited := make(map[p332edge]int)

	for _, tkt := range tickets {
		from := tkt[0]
		to := tkt[1]

		paths := graph[from]
		paths = append(paths, to)
		graph[from] = paths

		// update edge
		edge := p332edge{from, to}
		visited[edge]++
	}

	// sort paths
	for k, v := range graph {
		sort.Strings(v)
		graph[k] = v
	}

	currpath := []string{"JFK"}

	bestpath, _ := dfspaths(len(tickets)+1, graph, currpath, visited)
	fmt.Printf("bestpath: %v\n", bestpath)
	return bestpath
}

func dfspaths(explen int, graph map[string][]string, currpath []string, visited map[p332edge]int) ([]string, bool) {
	currnode := currpath[len(currpath)-1]
	fmt.Printf("visiting node : %s\n", currnode)

	if len(currpath) == explen {
		return currpath, true
	}

	nextnodes := graph[currnode]
	for _, next := range nextnodes {
		edge := p332edge{from: currnode, to: next}
		if visited[edge] == 0 {
			continue
		}

		visited[edge]--

		newpath := make([]string, len(currpath))
		copy(newpath, currpath)
		newpath = append(newpath, next)

		// var alldone bool
		finalpath, alldone := dfspaths(explen, graph, newpath, visited)
		visited[edge]++

		if alldone {
			return finalpath, true
		}
	}

	return currpath, false
}

type p332edge struct {
	from string
	to   string
}

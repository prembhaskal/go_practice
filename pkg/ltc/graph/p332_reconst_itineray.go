package graph

import (
	"fmt"
	"sort"
)

// TODO - do using https://www.geeksforgeeks.org/hierholzers-algorithm-directed-graph/ algorithm.
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

	allpaths := make([][]string, 0)
	currpath := []string{"JFK"}

	allpaths, _ = dfspaths(len(tickets)+1, graph, currpath, allpaths, visited)
	fmt.Printf("allpaths: %v\n", allpaths)

	validpaths := make([][]string, 0)
	for _, path := range allpaths {
		if len(path) == len(tickets)+1 {
			validpaths = append(validpaths, path)
		}
	}

	sort.Slice(validpaths, func(i, j int) bool {
		path1 := validpaths[i]
		path2 := validpaths[j]

		for k := 0; k < len(path1); k++ {
			if path1[k] < path2[k] {
				return true
			} else if path1[k] > path2[k] {
				return false
			}
		}

		return false
	})

	fmt.Printf("sorted valid paths: %v\n", validpaths)

	return validpaths[0]
}

func dfspaths(explen int, graph map[string][]string, currpath []string, allpaths [][]string, visited map[p332edge]int) ([][]string, bool) {
	currnode := currpath[len(currpath)-1]
	fmt.Printf("visiting node : %s\n", currnode)

	if len(currpath) == explen {
		newpath := make([]string, len(currpath))
		copy(newpath, currpath)
		allpaths = append(allpaths, newpath)
		return allpaths, true
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

		var alldone bool
		allpaths, alldone = dfspaths(explen, graph, newpath, allpaths, visited)
		visited[edge]++

		if alldone {
			return allpaths, true
		}
	}

	// if allpathsdone {
	// 	newpath := make([]string, len(currpath))
	// 	copy(newpath, currpath)
	// 	allpaths = append(allpaths, newpath)
	// }

	return allpaths, false
}

type p332edge struct {
	from string
	to   string
}

package graph

import (
	"fmt"
	"sort"
)

func findItinerary(tickets [][]string) []string {
	//     JFK -> [ATL|cross,  SFO|cross]

	graph := make(map[string][]string)
	for _, tkt := range tickets {
		from := tkt[0]
		to := tkt[1]

		paths := graph[from]
		paths = append(paths, to)
		graph[from] = paths
	}

	allpaths := make([][]string, 0)
	currpath := []string{"JFK"}
	visited := make(map[p332edge]bool)

	allpaths = dfspaths(graph, currpath, allpaths, visited)
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

func dfspaths(graph map[string][]string, currpath []string, allpaths [][]string, visited map[p332edge]bool) [][]string {
	currnode := currpath[len(currpath)-1]

	allpathsdone := true
	nextnodes := graph[currnode]
	for _, next := range nextnodes {
		edge := p332edge{from: currnode, to: next}
		if visited[edge] {
			continue
		}

		allpathsdone = false
		visited[edge] = true

		newpath := make([]string, len(currpath))
		copy(newpath, currpath)
		newpath = append(newpath, next)

		allpaths = dfspaths(graph, newpath, allpaths, visited)
		visited[edge] = false
	}

	if allpathsdone {
		newpath := make([]string, len(currpath))
		copy(newpath, currpath)
		allpaths = append(allpaths, newpath)
	}

	return allpaths
}

type p332edge struct {
	from string
	to   string
}

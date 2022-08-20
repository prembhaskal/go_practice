package queue_stack

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	return cloneGraphDFSRec(node)
}

func cloneGraphDFSRec(node *Node) *Node {
	visited := make([]bool, 101)
	nodecopies := make([]*Node, 101)
	if node == nil {
		return nil
	}

	nodecp := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0),
	}
	nodecopies[nodecp.Val] = nodecp

	cloneGraphRec(node, visited, nodecopies)
	return nodecp
}

func cloneGraphRec(curr *Node, visited []bool, nodecopies []*Node) {
	if visited[curr.Val] {
		return
	}
	visited[curr.Val] = true
	currcp := nodecopies[curr.Val]
	for _, neigh := range curr.Neighbors {
		neighcp := nodecopies[neigh.Val]
		if neighcp == nil {
			neighcp = &Node{Val: neigh.Val, Neighbors: make([]*Node, 0)}
			nodecopies[neighcp.Val] = neighcp
		}
		currcp.Neighbors = append(currcp.Neighbors, neighcp)
		cloneGraphRec(neigh, visited, nodecopies)
	}
}

func cloneGraphDFSStack(node *Node) *Node {
	visited := make([]bool, 101)
	copynodes := make([]*Node, 101)

	stk := newnodestack()
	// visited[node.Val] = true

	if node == nil {
		return nil
	}

	stk.push(node)
	nodecopy := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0),
	}
	copynodes[nodecopy.Val] = nodecopy

	for !stk.isempty() {
		curr := stk.pop()
		if visited[curr.Val] {
			continue
		}
		visited[curr.Val] = true
		currcopy := copynodes[curr.Val]

		// fmt.Printf("copying node: %d, orig: %p, copy: %p\n", curr.Val, curr, currcopy )
		// fmt.Printf("original node neighbours: %v\n", curr.Neighbors)

		for _, neigh := range curr.Neighbors {
			neighcopy := copynodes[neigh.Val]
			if neighcopy == nil {
				neighcopy = &Node{
					Val:       neigh.Val,
					Neighbors: make([]*Node, 0),
				}
			}
			copynodes[neighcopy.Val] = neighcopy
			// fmt.Printf("copied node: %p to copy: %p\n", neigh, neighcopy)

			currcopy.Neighbors = append(currcopy.Neighbors, neighcopy)
			// neighcopy.Neighbors = append(neighcopy.Neighbors, currcopy)

			// fmt.Printf("adding neighbour: %p to: %p\n", neighcopy,  currcopy)
			stk.push(neigh)
		}
	}

	return nodecopy
}

type nodestack struct {
	arr []*Node
}

func newnodestack() *nodestack {
	return &nodestack{
		arr: make([]*Node, 0),
	}
}

func (s *nodestack) push(node *Node) {
	s.arr = append(s.arr, node)
}

func (s *nodestack) pop() *Node {
	if s.isempty() {
		return nil
	}

	node := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return node
}

func (s *nodestack) isempty() bool {
	return len(s.arr) == 0
}

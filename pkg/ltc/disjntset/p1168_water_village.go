package disjntset

import "sort"

// model as graph problems.
// underground water source as (n+1)th node.
// house numbers are numbered from 1 to n.
// sovle MST and return its cost.
func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
	// edges
	edges := make([]p1168edge, 0)

	// pipe - house1, house2, dist.
	for _, pipe := range pipes {
		h1 := pipe[0]
		h2 := pipe[1]
		cst := pipe[2]

		edge := p1168edge{h1, h2, cst}
		edges = append(edges, edge)
	}

	// for each house, add a edge to well. (its number is n+1)
	for i := 1; i <= n; i++ {
		edge := p1168edge{i, n + 1, wells[i-1]}
		edges = append(edges, edge)
	}

	// run kruskal algorithm.
	// sort the edges
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].cost < edges[j].cost
	})

	ufind := newp1168ufind(n + 2) // n+1 th for the well, +1 because of 1 based numbering.

	totalcost := 0
	for _, edge := range edges {
		// check if edge makes a cycle in
		a := edge.from
		b := edge.to
		if ufind.findroot(a) != ufind.findroot(b) {
			ufind.union(a, b)
			totalcost = totalcost + edge.cost
		}
	}

	return totalcost
}

type p1168edge struct {
	from int
	to   int
	cost int
}

type p1168ufind struct {
	root  []int
	rank  []int
	count int
}

func newp1168ufind(n int) *p1168ufind {
	root := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
	}

	return &p1168ufind{
		root:  root,
		rank:  make([]int, n),
		count: n,
	}
}

func (u *p1168ufind) findroot(x int) int {
	if x == u.root[x] {
		return x
	}
	rt := u.findroot(u.root[x])
	u.root[x] = rt
	return rt
}

func (u *p1168ufind) union(x, y int) {
	rx := u.findroot(x)
	ry := u.findroot(y)
	if rx != ry {
		if u.rank[rx] > u.rank[ry] {
			u.root[ry] = u.root[rx]
		} else if u.rank[ry] > u.rank[rx] {
			u.root[rx] = u.root[ry]
		} else {
			u.root[ry] = u.root[rx]
			u.rank[rx] = u.rank[rx] + 1
		}

		u.count--
	}
}

package disjntset

func validPath(n int, edges [][]int, source int, destination int) bool {
	ufind := newp1971ufind(n)
	for _, edge := range edges {
		ufind.union(edge[0], edge[1])
	}

	return ufind.findroot(source) == ufind.findroot(destination)
}

type p1971ufind struct {
	root []int
	rank []int
}

func newp1971ufind(n int) *p1971ufind {
	root := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
	}

	return &p1971ufind{
		root: root,
		rank: make([]int, n),
	}
}

func (u *p1971ufind) findroot(x int) int {
	if x == u.root[x] {
		return x
	}

	rt := u.findroot(u.root[x])
	u.root[x] = rt
	return rt
}

func (u *p1971ufind) union(x, y int) {
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
	}
}

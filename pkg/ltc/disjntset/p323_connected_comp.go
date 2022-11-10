package disjntset

func countComponents(n int, edges [][]int) int {
	ufind := newP323UF(n)
	for _, edge := range edges {
		ufind.union(edge[0], edge[1])
	}
	return ufind.count
}

type p323UF struct {
	root  []int
	rank  []int
	count int
}

func newP323UF(n int) *p323UF {
	root := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
	}

	return &p323UF{
		root:  root,
		rank:  make([]int, n),
		count: n,
	}
}

func (u *p323UF) findRoot(x int) int {
	if x == u.root[x] {
		return x
	}
	rt := u.findRoot(u.root[x])
	u.root[x] = rt
	return rt
}

func (u *p323UF) union(x, y int) {
	rx := u.findRoot(x)
	ry := u.findRoot(y)
	if rx != ry {
		if u.rank[rx] > u.rank[ry] {
			u.root[ry] = rx
		} else if u.rank[ry] > u.rank[rx] {
			u.root[rx] = ry
		} else {
			u.root[ry] = rx
			u.rank[rx]++
		}

		u.count--

	}
}

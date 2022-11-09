package disjntset

func validTree(n int, edges [][]int) bool {
	ufind := newP261UF(n)
	for _, edge := range edges {
		a := edge[0]
		b := edge[1]
		alreadyconnected := ufind.union(a, b)
		if alreadyconnected {
			return false
		}
	}

	if ufind.count != 1 {
		return false
	}

	return true
}

type P261UF struct {
	root  []int
	rank  []int
	count int
}

func newP261UF(n int) *P261UF {
	root := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
	}

	return &P261UF{
		root:  root,
		rank:  make([]int, n),
		count: n,
	}
}

func (u *P261UF) findroot(x int) int {
	if x == u.root[x] {
		return x
	}
	rt := u.findroot(u.root[x])
	u.root[x] = rt
	return rt
}

// union, also return true if already connected.
func (u *P261UF) union(x, y int) bool {
	rx := u.findroot(x)
	ry := u.findroot(y)
	if rx == ry {
		return true
	}
	if rx != ry {
		if u.rank[rx] > u.rank[ry] {
			u.root[ry] = rx
		} else if u.rank[rx] < u.rank[ry] {
			u.root[rx] = ry
		} else {
			u.root[ry] = rx
			u.rank[x] = u.rank[x] + 1
		}
		u.count--
	}
	return false
}

package disjntset

import "sort"

func earliestAcq(logs [][]int, n int) int {
	// sort by time
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	ufind := newP1101UF(n)

	earliesttime := -1
	for _, log := range logs {
		tm := log[0]
		x := log[1]
		y := log[2]
		ufind.union(x, y)
		if ufind.count == 1 {
			if earliesttime == -1 {
				earliesttime = tm
			}
		} else { // some new edge had new vertices not part of existing conn. components.
			earliesttime = -1
		}
	}

	return earliesttime
}

type p1101UF struct {
	root  []int
	rank  []int
	count int
}

func newP1101UF(n int) *p1101UF {
	root := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
	}
	return &p1101UF{
		root:  root,
		rank:  make([]int, n),
		count: n,
	}
}

func (u *p1101UF) findroot(x int) int {
	if x == u.root[x] {
		return x
	}
	rt := u.findroot(u.root[x])
	u.root[x] = rt
	return rt
}

func (u *p1101UF) union(x, y int) {
	rx := u.findroot(x)
	ry := u.findroot(y)
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

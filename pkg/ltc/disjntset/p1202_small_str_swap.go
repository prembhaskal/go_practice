package disjntset

import "sort"

func smallestStringWithSwaps(s string, pairs [][]int) string {
	rns := []rune(s)
	ufind := newp1202UF(len(s))
	for _, pair := range pairs {
		ufind.union(pair[0], pair[1])
	}

	rootList := make(map[int][]int, 0)
	for i := 0; i < len(s); i++ {
		rt := ufind.findroot(i)
		list := rootList[rt]
		list = append(list, i)
		rootList[rt] = list
	}

	// new string
	newrn := make([]rune, len(rns))

	// sort each root list based on char present at that index.
	for _, list := range rootList {
		rnlist := make([]rune, len(list))
		i := 0
		// get chars at the root
		for _, v := range list {
			rnlist[i] = rns[v]
			i++
		}

		// sort chars
		sort.Slice(rnlist, func(i, j int) bool {
			return rnlist[i] < rnlist[j]
		})

		// put chars back in new strings, based on location from list.
		j := 0
		for _, v := range list {
			newrn[v] = rnlist[j]
			j++
		}
	}

	return string(newrn)
}

type p1202UF struct {
	root  []int
	rank  []int
	count int
}

func newp1202UF(n int) *p1202UF {
	root := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
	}
	return &p1202UF{
		root:  root,
		rank:  make([]int, n),
		count: n,
	}
}

func (u *p1202UF) findroot(x int) int {
	if x == u.root[x] {
		return x
	}
	rt := u.findroot(u.root[x])
	u.root[x] = rt
	return rt
}

func (u *p1202UF) union(x, y int) {
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

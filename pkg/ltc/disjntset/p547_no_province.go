package disjntset

func findCircleNum(isConnected [][]int) int {
	// create a union-find ds, by rank, path compressed
	// for each connection, update union find.  - O(n^2) //  can we optimize this??
	// since input is matrix, i guess nothing we can, but to traverse
	// how to find all provinces ??
	// for each node, find root, capture uniq roots  - O(n)
	n := len(isConnected)
	ufind := newp547UF(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				ufind.union(i, j)
			}
		}
	}

	// fmt.Printf("ufind: %v\n", ufind.root)

	//     uniqRoot := make(map[int]bool, 0)
	//     cnt := 0
	//     for i :=0; i < n; i++ {
	//         rootI := ufind.findRoot(i)
	//         if !uniqRoot[rootI] {
	//             uniqRoot[rootI] = true
	//             cnt++
	//         }
	//     }

	//     return cnt

	return ufind.getCount()
}

type p547UF struct {
	root  []int
	rank  []int
	count int
}

func newp547UF(n int) *p547UF {
	root := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
	}
	return &p547UF{
		root:  root,
		rank:  make([]int, n),
		count: n,
	}
}

func (u *p547UF) findRoot(x int) int {
	if x == u.root[x] {
		return x
	}
	rt := u.findRoot(u.root[x])
	u.root[x] = rt
	return rt
}

func (u *p547UF) union(x, y int) {
	rootX := u.findRoot(x)
	rootY := u.findRoot(y)

	// update roots and their ranks based on ranks.
	if rootX != rootY {
		if u.rank[rootX] > u.rank[rootY] {
			u.root[rootY] = rootX
		} else if u.rank[rootX] < u.rank[rootY] {
			u.root[rootX] = rootY
		} else {
			// merge y into x, increase rank of x.
			u.root[rootY] = rootX
			u.rank[rootX] = u.rank[rootX] + 1
		}
		u.count--
	}
}

func (u *p547UF) getCount() int {
	return u.count
}

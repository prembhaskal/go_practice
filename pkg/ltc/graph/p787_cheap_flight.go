package graph

import "fmt"

/*
There are n cities connected by some number of flights.
You are given an array flights where flights[i] = [fromi, toi, pricei]
indicates that there is a flight from city fromi to city toi with cost pricei.

You are also given three integers src, dst, and k,
return the cheapest price from src to dst with at most k stops.
If there is no such route, return -1.
*/
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// return findCheapestPriceCLRS(n, flights, src, dst, k)
	return findCheapestPriceDP(n, flights, src, dst, k)
}

func findCheapestPriceDP(n int, flights [][]int, src int, dst int, k int) int {
	// DP[k][v] = min{ for u in 0:n DP[k-1][u] + W(u,v), DP[k][v]}

	adj := make(map[int][]edge787)
	for _, flight := range flights {
		src := flight[0]
		dst := flight[1]
		price := flight[2]
		adj[src] = append(adj[src], edge787{dst, price})
	}

	inf := 100000 // 10^5
	dp := make([][]int, 0)

	k = k + 1
	for i := 0; i <= k; i++ {
		row := make([]int, n)
		dp = append(dp, row)
		for j := 0; j < n; j++ {
			dp[i][j] = inf
		}
	}

	dp[0][src] = 0
	for i := 1; i <= k; i++ {
		for u := 0; u < n; u++ {
			if dp[i][u] > dp[i-1][u] {
				dp[i][u] = dp[i-1][u]
			}
			for _, to := range adj[u] {
				v := to.to
				if dp[i][v] > dp[i-1][u]+to.price {
					dp[i][v] = dp[i-1][u] + to.price
					fmt.Printf("%d iter k:%d, relaxed edge:%d-%d, new price: %d\n", i, k, u, v, dp[i][v])
				}
			}
		}
	}

	if dp[k][dst] == inf {
		return -1
	}

	return dp[k][dst]
}

type pair787 struct {
	from int
	to   int
}

func findCheapestPriceCLRS(n int, flights [][]int, src int, dst int, k int) int {
	// bellman ford

	// type Vertex {num int, d int}
	// for i = 0 to n - 1
	// for each edge (u,v) E Graph
	// graph = map[int][]Edge{to int, dst int}
	// relax(u,v)
	//    if v.d > u.d + W(u,v)
	//        v.d = u.d + W(u,v)

	// init source
	inf := 100000
	vertices := make([]*vtx787, n)
	for i := 0; i < n; i++ {
		vertices[i] = &vtx787{i, inf, inf}
	}

	vertices[src].d = 0

	adj := make(map[int][]edge787)
	for _, flight := range flights {
		src := flight[0]
		dst := flight[1]
		price := flight[2]
		adj[src] = append(adj[src], edge787{dst, price})
	}

	// edges = stops + 1
	k = k + 1

	// run k times at max
	for i := 1; i <= k; i++ {
		for un := 0; un < n; un++ {
			u := vertices[un]
			for _, to := range adj[un] {
				v := vertices[to.to]
				if v.nd > u.d+to.price {
					v.nd = u.d + to.price
					// fmt.Printf("k=%d, relaxed edge: %d-%d, with new dist: %d\n", i, u.num, v.num, v.nd)
				}
			}
		}
		// update d with nd
		for _, vtx := range vertices {
			vtx.d = vtx.nd
		}
	}

	if vertices[dst].d == inf {
		return -1
	}

	return vertices[dst].d
}

type edge787 struct {
	to    int
	price int
}

type vtx787 struct {
	num int
	d   int
	nd  int
}

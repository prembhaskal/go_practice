package graph

import "fmt"

func findMinHeightTrees(n int, edges [][]int) []int {
	adjlist := make([][]int, n)
	degree := make([]int, n)
	for _, edge := range edges {
		a := edge[0]
		b := edge[1]
		adjlist[a] = append(adjlist[a], b)
		adjlist[b] = append(adjlist[b], a)
		degree[a]++
		degree[b]++
	}

	vertices := make([]*vtx310, n)
	for i := 0; i < n; i++ {
		vertices[i] = &vtx310{i, 0}
	}

	q := newqueue310()
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			q.add(vertices[i])
			visited[i] = true
		}
	}

	// q.print()

	minht := 0

	for q.size() > 0 {
		item := q.poll()
		for _, nextitem := range adjlist[item.val] {
			next := vertices[nextitem]
			if !visited[next.val] {
				degree[next.val]--
				if degree[next.val] == 1 {
					visited[next.val] = true
					next.ht = item.ht + 1
					q.add(next)
					if next.ht > minht {
						minht = next.ht
					}
				}
			}
		}
	}

	mht := make([]int, 0)
	for i := 0; i < n; i++ {
		if vertices[i].ht == minht {
			mht = append(mht, vertices[i].val)
		}
	}

	return mht
}

type vtx310 struct {
	val int
	ht  int
}

type queue310 struct {
	ar []*vtx310
}

func newqueue310() *queue310 {
	return &queue310{
		ar: make([]*vtx310, 0),
	}
}

func (q *queue310) size() int {
	return len(q.ar)
}

func (q *queue310) add(item *vtx310) {
	q.ar = append(q.ar, item)
}

func (q *queue310) poll() *vtx310 {
	if q.size() == 0 {
		panic("empty queue")
	}
	item := q.ar[0]
	q.ar = q.ar[1:]
	return item
}

func (q *queue310) print() {
	for i := 0; i < len(q.ar); i++ {
		fmt.Printf("-> %d ", q.ar[i].val)
	}
	fmt.Println()
}

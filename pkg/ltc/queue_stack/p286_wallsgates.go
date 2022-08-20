package queue_stack

func wallsAndGates(rooms [][]int) {
	iq := newintq286()
	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[i]); j++ {
			if rooms[i][j] == 0 {
				iq.enq(newnode286(i, j))
				// doBFS(rooms, iq)
			}
		}
	}

	doBFSWalls(rooms, iq)
}

func doBFSWalls(rooms [][]int, iq *intq286) {
	for !iq.isempty() {
		item := iq.deq()
		// fill the neighbours
		dist := rooms[item.i][item.j]
		addNeighbour(rooms, iq, dist+1, item.i, item.j+1)
		addNeighbour(rooms, iq, dist+1, item.i, item.j-1)
		addNeighbour(rooms, iq, dist+1, item.i+1, item.j)
		addNeighbour(rooms, iq, dist+1, item.i-1, item.j)
	}
}

func addNeighbour(rooms [][]int, iq *intq286, dist, i, j int) {
	if j >= len(rooms[0]) || j < 0 {
		return
	}
	if i >= len(rooms) || i < 0 {
		return
	}
	if rooms[i][j] == 2147483647 { // update only non-visited nodes.
		rooms[i][j] = dist
		iq.enq(newnode286(i, j))
	}
}

type node286 struct {
	i int
	j int
}

func newnode286(i, j int) *node286 {
	return &node286{i, j}
}

type intq286 struct {
	ar []*node286
}

func newintq286() *intq286 {
	return &intq286{
		ar: make([]*node286, 0),
	}
}

func (q *intq286) enq(item *node286) {
	q.ar = append(q.ar, item)
}

func (q *intq286) deq() *node286 {
	if q.isempty() {
		return nil
	}
	val := q.ar[0]
	q.ar = q.ar[1:]
	return val
}

func (q *intq286) isempty() bool {
	return len(q.ar) == 0
}

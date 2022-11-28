package graph

// TODO - add the necessary stuff to compile below stuff
// TODO - solve using binary search approach too.
func minimumEffortPath(heights [][]int) int {
	return -1
	// rows := len(heights)
	// cols := len(heights[0])

	// minq := newminq();

	// dirs := [][]int{{-1,0}, {1,0}, {0,1}, {0,-1}}

	// vertices := make([][]Vtx, 0) // m x n size

	// for minq.size() > 0 {
	// 	u := minq.extractmin()
	// 	visited[{u.row, u.col}] = true

	// 	currht := heights[u.row][u.col]

	// 	for _, dir := range dirs {
	// 		nr := u.row + dir[0]
	// 		nc := u.col + dir[1]
	// 		if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
	// 			continue
	// 		}

	// 		if visited[{nr, nc}] {
	// 			continue
	// 		}

	// 		if minq.contains({nr, nc}) {
	// 			minq.insert(v)
	// 		}

	// 		v := vertices[nr][nc]
	// 		if v.d > u.d + abs(heights[nr][nc] - currht) {
	// 			v.d = u.d + abs(heights[nr][nc] - currht)
	// 			minq.decreasekey({nr, nc}, v.d)
	// 		}

	// 	}

	// }

	// return vertices[rows-1][cols-1].d
}

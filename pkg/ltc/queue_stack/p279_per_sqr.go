package queue_stack

func numSquares(n int) int {
	// bfs approach
	// start at 0 and move to next by adding a perfect square
	sqrs := make([]int, 101)
	for i := 1; i < 101; i++ {
		sqrs[i] = i * i
	}

	visited := make([]bool, n+1)
	ways := make([]int, n+1)

	iq := newintq279()
	iq.in(0)
	ways[0] = 0
	visited[0] = true // we mark visited after adding to queue to track the right number of ways without comaparison

	for !iq.isempty() {
		item := iq.out()
		// if visited[item] {
		//     continue
		// }
		// fmt.Printf("got from queue: %d\n", item)

		if item == n {
			return ways[item]
		}

		// visited[item] = true
		// find next neighour
		for _, sq := range sqrs {
			newval := item + sq
			if newval > n {
				continue
			}
			if visited[newval] {
				continue
			}
			// fmt.Printf("check ways[%d]: %d\n", newval, ways[newval])
			// if we have found a shorter way, check next
			//             if ways[newval] < ways[item] + 1 {
			//                 continue
			//             }

			visited[newval] = true
			iq.in(newval)
			ways[newval] = ways[item] + 1
			// fmt.Printf("adding to queue: %d, hops: %d\n", newval, ways[newval])
		}

	}

	panic("wrong..should always find answer")
}

type int279 struct {
	arr []int
}

func newintq279() *int279 {
	return &int279{
		arr: make([]int, 0),
	}
}

func (i *int279) in(val int) {
	i.arr = append(i.arr, val)
}

func (i *int279) out() int {
	if i.isempty() {
		return -1
	}
	val := i.arr[0]
	i.arr = i.arr[1:]
	return val
}

func (i *int279) isempty() bool {
	return len(i.arr) == 0
}

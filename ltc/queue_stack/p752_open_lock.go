package queue_stack

import "strconv"

type locknode struct {
	lock []int
	ln   int
}

func openLock(deadends []string, target string) int {
	iq := newintq()
	visited := make([]bool, 10000)
	for _, dead := range deadends {
		deadint, _ := strconv.Atoi(dead)
		visited[deadint] = true
	}

	trg, _ := strconv.Atoi(target)
	steps := -1

	lock := []int{0, 0, 0, 0}
	iq.add(locknode{lock: lock, ln: 0})

	for !iq.isempty() {
		item, _ := iq.poll()
		itemval := arrToInt(item.lock)
		if visited[itemval] {
			continue
		}
		if itemval == trg {
			return item.ln
		}
		// fmt.Printf("visiting node %d: steps: %d\n", itemval, item.ln)

		visited[itemval] = true

		addNext(iq, item)
	}

	return steps
}

func arrToInt(arr []int) int {
	// 0 digit is MSB
	total := 0
	for _, val := range arr {
		total = total*10 + val
	}
	return total
}

func addNext(iq *intq, curr locknode) {
	lock := curr.lock
	for i := 0; i < 4; i++ {
		next := (lock[i] + 1) % 10
		nextlock := []int{lock[0], lock[1], lock[2], lock[3]}
		nextlock[i] = next
		iq.add(locknode{lock: nextlock, ln: curr.ln + 1})

		prev := (lock[i] + 9) % 10
		prevlock := []int{lock[0], lock[1], lock[2], lock[3]}
		prevlock[i] = prev
		iq.add(locknode{lock: prevlock, ln: curr.ln + 1})
	}

}

type intq struct {
	ar []locknode
}

func newintq() *intq {
	return &intq{
		ar: make([]locknode, 0),
	}
}

func (i *intq) add(item locknode) {
	i.ar = append(i.ar, item)
}

func (i *intq) poll() (locknode, bool) {
	if i.isempty() {
		return locknode{}, false
	}
	item := i.ar[0]
	i.ar = i.ar[1:]
	return item, true
}

func (i *intq) isempty() bool {
	return len(i.ar) == 0
}

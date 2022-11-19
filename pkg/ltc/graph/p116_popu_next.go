package graph

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node116 struct {
	Val   int
	Left  *Node116
	Right *Node116
	Next  *Node116
}

func connect(root *Node116) *Node116 {
	// BFS with level chekc

	vq := newvqueue()
	rootver := vertex116{root, 0}
	if root != nil {
		vq.add(rootver)
	}

	for vq.size() > 0 {
		currver := vq.poll()

		if vq.size() > 0 {
			nextver := vq.peek()
			if currver.level == nextver.level {
				currver.node.Next = nextver.node
			}
		}

		// add next node to queue
		if currver.node.Left != nil {
			vq.add(vertex116{currver.node.Left, currver.level + 1})
		}
		if currver.node.Right != nil {
			vq.add(vertex116{currver.node.Right, currver.level + 1})
		}
	}

	return root
}

type vertex116 struct {
	node  *Node116
	level int
}

type vqueue116 struct {
	ar []vertex116
}

func newvqueue() *vqueue116 {
	return &vqueue116{
		ar: make([]vertex116, 0),
	}
}

func (q *vqueue116) size() int {
	return len(q.ar)
}

func (q *vqueue116) add(item vertex116) {
	q.ar = append(q.ar, item)
}

func (q *vqueue116) peek() vertex116 {
	if q.size() == 0 {
		panic("empty queue")
	}
	return q.ar[0]
}

func (q *vqueue116) poll() vertex116 {
	if q.size() == 0 {
		panic("empty queue")
	}
	item := q.ar[0]
	q.ar = q.ar[1:]
	return item
}

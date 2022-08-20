package bintree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	lq := newlevelorderqueue()

	lnodes := make([][]int, 0)
	if root == nil {
		return lnodes
	}
	lq.enq(&levelOrderNode{root, 0})

	for !lq.isempty() {
		curr := lq.deq().(*levelOrderNode)
		// curr.level
		// lnodes[curr.level] <-- append to this.
		if len(lnodes) < curr.level+1 {
			lnodes = append(lnodes, make([]int, 0))
		}
		lnodes[curr.level] = append(lnodes[curr.level], curr.node.Val)

		if curr.node.Left != nil {
			lq.enq(&levelOrderNode{curr.node.Left, curr.level + 1})
		}
		if curr.node.Right != nil {
			lq.enq(&levelOrderNode{curr.node.Right, curr.level + 1})
		}
	}

	return lnodes
}

type levelOrderNode struct {
	node  *TreeNode
	level int
}

type t interface{}

type levelorderqueue struct {
	ar []t
}

func newlevelorderqueue() *levelorderqueue {
	return &levelorderqueue{
		ar: make([]t, 0),
	}
}

func (q *levelorderqueue) enq(item t) {
	q.ar = append(q.ar, item)
}

func (q *levelorderqueue) deq() t {
	var item t
	if q.isempty() {
		return item
	}

	item = q.ar[0]
	q.ar = q.ar[1:]
	return item
}

func (q *levelorderqueue) isempty() bool {
	return len(q.ar) == 0
}

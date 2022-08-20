package bintree

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	// return connectusingqueue(root)
	return connectdfs(root)
}

// work for full tree only.
func connectdfs(curr *Node) *Node {
	if curr == nil {
		return nil
	}

	/*
	   connect curr right element to left of current element's Next.
	   for this it is needed that we set Next element of curr before going to next level.

	   curr.left.Next = curr.right
	   curr.right.Next = curr.Next == null ? null : curr.Next.left
	   connectdfs(curr.left)
	   connectdfs(curr.right)
	*/

	if curr.Left != nil {
		curr.Left.Next = curr.Right
	}
	if curr.Right != nil {
		var nextleft *Node
		if curr.Next != nil {
			nextleft = curr.Next.Left
		}
		curr.Right.Next = nextleft
	}
	connectdfs(curr.Left)
	connectdfs(curr.Right)
	return curr
}

func connectusingqueue(root *Node) *Node {
	if root == nil {
		return nil
	}

	nq := newnodequeue()

	rootNode := newlevelnode(root, 0)

	nq.enque(rootNode)

	// prev := nil
	for !nq.isempty() {
		nd := nq.deque().(*levelnode)
		// fmt.Printf("node: %d, level: %d\n", nd.Val, nd.level)
		nextnodei := nq.peek()
		if nextnodei != nil {
			nextnode := nextnodei.(*levelnode)
			if nextnode.level != nd.level {
				nd.Next = nil
			} else {
				nd.Next = nextnode.Node
			}
		} else {
			nd.Next = nil
		}
		if nd.Left != nil {
			nq.enque(newlevelnode(nd.Left, nd.level+1))
		}
		if nd.Right != nil {
			nq.enque(newlevelnode(nd.Right, nd.level+1))
		}

		// prev := nd.node
	}

	return root
}

type levelnode struct {
	*Node
	level int
}

func newlevelnode(node *Node, level int) *levelnode {
	return &levelnode{
		Node:  node,
		level: level,
	}
}

type t1 interface{}

type nodequeue struct {
	ar []t1
}

func newnodequeue() *nodequeue {
	return &nodequeue{
		ar: make([]t1, 0),
	}
}

// insert at tail
func (n *nodequeue) enque(elem t1) {
	n.ar = append(n.ar, elem)
}

// remove from head
func (n *nodequeue) deque() t1 {
	var elem t1
	if n.isempty() {
		return elem
	}

	elem = n.ar[0]
	n.ar = n.ar[1:]
	return elem
}

func (n *nodequeue) peek() t1 {
	var elem t1
	if n.isempty() {
		return elem
	}
	return n.ar[0]
}

func (n *nodequeue) isempty() bool {
	return len(n.ar) == 0
}

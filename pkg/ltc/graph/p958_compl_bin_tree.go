package graph

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	mq := newmyqueue()
	mq.add(root)
	seennil := false
	for mq.size() > 0 {
		item := mq.poll()
		if item == nil {
			seennil = true
			continue
		}
		if seennil { // got non nil item after previous nil
			return false
		}
		// fmt.Printf("poll val: %d\n", item.Val)
		mq.add(item.Left)
		mq.add(item.Right)
	}
	return true
}

type myqueue struct {
	ar []*TreeNode
}

func newmyqueue() *myqueue {
	mq := &myqueue{
		ar: make([]*TreeNode, 0),
	}
	return mq
}

func (q *myqueue) add(item *TreeNode) {
	q.ar = append(q.ar, item)
}

func (q *myqueue) poll() *TreeNode {
	if q.size() == 0 {
		panic("poll on empty queue")
	}
	item := q.ar[0]
	q.ar = q.ar[1:len(q.ar)]
	return item
}

func (q *myqueue) size() int {
	return len(q.ar)
}

package bintree

type NumArray struct {
	root *segNode
	n    int
}

func Constructor(nums []int) NumArray {
	segNode := createsegtree(nums)
	return NumArray{
		root: segNode,
		n:    len(nums),
	}
}

func (this *NumArray) Update(index int, val int) {
	updatetree(index, val, this.n-1, this.root)
}

func (this *NumArray) SumRange(left int, right int) int {
	return querytree(left, right, 0, this.n-1, this.root)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

// TODO - current implementation runs slow - > 900ms, whereas same implementation in java runs in 112ms , find root cause.
// https://leetcode.com/problems/range-sum-query-mutable/discuss/75724/17-ms-Java-solution-with-segment-tree  Java impl.
//
//	TODO 1 - implement with a array instaed of linked nodes.
//	TODO
//
// segment tree
type segNode struct {
	left  *segNode
	right *segNode
	val   int
}

func newsegnode(left, right *segNode, val int) *segNode {
	return &segNode{
		left:  left,
		right: right,
		val:   val,
	}
}

func createsegtree(nums []int) *segNode {
	return createtree(0, len(nums)-1, nums)
}

func createtree(left, right int, nums []int) *segNode {
	if left == right {
		return newsegnode(nil, nil, nums[left])
	}

	mid := (left + right) / 2
	leftNode := createtree(left, mid, nums)
	rightNode := createtree(mid+1, right, nums)

	return newsegnode(leftNode, rightNode, leftNode.val+rightNode.val)
}

func updatetree(idx, newval int, n int, root *segNode) {
	update(0, n, idx, newval, root)
}

func update(left, right, idx, newval int, root *segNode) {
	// return if out of bounds
	if idx < left || idx > right {
		// fmt.Printf("returning - left: %d, right: %d, idx: %d\n", left, right, idx)
		return
	}

	// update leaf
	if left == right {
		// fmt.Printf("update leaf with idx: %d with new val: %d\n", left, newval)
		root.val = newval
		return
	}

	mid := (left + right) / 2
	// update child nodes
	update(left, mid, idx, newval, root.left)
	update(mid+1, right, idx, newval, root.right)

	// update self
	// oldval := root.val
	root.val = root.left.val + root.right.val
	// fmt.Printf("updated node[%d:%d] with new value: %d, old value: %d\n", left, right, root.val, oldval)
}

func querytree1(ql, qr, left, right int, root *segNode) int { // variation where the range ql, qr are not changed, only the conditions are changed
	// fmt.Printf("query ql: %d, qr: %d, left: %d, right: %d\n", ql, qr, left, right)
	// left .... right  ql .... qr
	//  ql ... qr left .... right
	// non overlap
	if ql > right || qr < left {
		// fmt.Printf("non overlap, returning\n")
		return 0
	}

	// ql ... left ... right ....qr
	// full within
	if ql <= left && qr >= right {
		return root.val
	}

	// partial overlap
	mid := (left + right) / 2
	return querytree1(ql, qr, left, mid, root.left) + querytree1(ql, qr, mid+1, right, root.right)
}

func querytree(ql, qr, left, right int, root *segNode) int {
	// fmt.Printf("query ql: %d, qr: %d, left: %d, right: %d\n", ql, qr, left, right)
	// left .... right  ql .... qr
	//  ql ... qr left .... right
	// non overlap
	// if ql > right || qr < left {
	// // fmt.Printf("non overlap, returning\n")
	// 	return 0
	// }

	// ql ... left ... right ....qr
	// full within
	if ql == left && qr == right {
		return root.val
	}

	// partial overlap
	mid := (left + right) / 2
	if qr <= mid { // only in left
		return querytree(ql, qr, left, mid, root.left)
	} else if ql >= mid+1 { // only in right
		return querytree(ql, qr, mid+1, right, root.right)
	} else {
		return querytree(ql, mid, left, mid, root.left) + querytree(mid+1, qr, mid+1, right, root.right)
	}
}

// public int sumRange(SegmentTreeNode root, int start, int end) {
//     if (root.end == end && root.start == start) {
//         return root.sum;
//     } else {
//         int mid = root.start + (root.end - root.start) / 2;
//         if (end <= mid) {
//             return sumRange(root.left, start, end);
//         } else if (start >= mid+1) {
//             return sumRange(root.right, start, end);
//         }  else {
//             return sumRange(root.right, mid+1, end) + sumRange(root.left, start, mid);
//         }
//     }
// }

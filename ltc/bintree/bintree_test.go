package bintree

import (
	"testing"
)

func TestTraverse(t *testing.T) {

	n8 := &TreeNode{8, nil, nil}
	n7 := &TreeNode{7, nil, nil}
	n6 := &TreeNode{6, n7, n8}

	n5 := &TreeNode{5, nil, nil}
	n4 := &TreeNode{4, nil, nil}
	n3 := &TreeNode{3, n4, n5}

	n2 := &TreeNode{2, n3, nil}
	n01 := &TreeNode{11, nil, nil}
	n1 := &TreeNode{1, n6, n01}
	n0 := &TreeNode{0, n1, n2}

	ir := inorderTraversal(n0)
	t.Logf("inorder :%v", ir)

	pr := postorderTraversal(n0)
	t.Logf("post order: %v", pr)

	//             0
	//       1              2
	//   nil    6        3      nil
	//       7    8   4    5

}

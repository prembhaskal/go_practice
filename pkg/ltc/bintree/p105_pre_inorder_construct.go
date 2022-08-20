package bintree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(inorder)
	posmap := make(map[int]int)
	for i, v := range inorder {
		posmap[v] = i
	}

	return buildTreePI(preorder, inorder, 0, n-1, 0, n-1, posmap)
}

func buildTreePI(po, in []int, ps, pe, is, ie int, inmap map[int]int) *TreeNode {
	if ps > pe || is > ie {
		return nil
	}

	// first element of pre-order is root
	root := po[ps]

	inpos := inmap[root]

	// left elements = end-start+1
	leftelems := (inpos - 1) - is + 1
	// skip ps and inpos , the root elements
	// end index formula = start + total - 1
	leftNode := buildTreePI(po, in, ps+1, ps+1+leftelems-1, is, inpos-1, inmap)
	rightNode := buildTreePI(po, in, ps+1+leftelems, pe, inpos+1, ie, inmap)
	return &TreeNode{
		Val:   root,
		Left:  leftNode,
		Right: rightNode,
	}
}

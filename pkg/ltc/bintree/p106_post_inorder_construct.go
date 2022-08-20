package bintree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree2(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	inmap := getposmap(inorder)
	return buildTree1(inorder, postorder, 0, n-1, 0, n-1, inmap)
}

func getposmap(ar []int) map[int]int {
	posmap := make(map[int]int)
	for i, v := range ar {
		posmap[v] = i
	}
	return posmap
}

func buildTree1(in, post []int, is, ie, ps, pe int, inmap map[int]int) *TreeNode {
	// fmt.Printf("build tree: is:%d, ie:%d, ps:%d,  pe:%d\n", is, ie, ps, pe)
	if is > ie || ps > pe {
		// fmt.Printf("base case\n")
		return nil
	}
	root := post[pe] // last element of post order is root
	// fmt.Printf("root is %d\n", root)
	// in-order-tree   [left-root-righ] - [parent-root] - [left-root-right]
	// find pos of root in in-order
	inroot := inmap[root]

	// inorder bounds for lefttree =>  is -- inroot-1
	// total elements (inroot-1) - is + 1 = inroot - is
	// postorder bounds for leftree => ps -- ps + (inroot - is)
	leftelems := inroot - is
	leftNode := buildTree1(in, post, is, inroot-1, ps, ps+leftelems-1, inmap)

	rightNode := buildTree1(in, post, inroot+1, ie, ps+leftelems, pe-1, inmap)

	return &TreeNode{
		Val:   root,
		Left:  leftNode,
		Right: rightNode,
	}
}

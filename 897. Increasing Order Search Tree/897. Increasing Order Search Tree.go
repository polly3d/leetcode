/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	s := []int{}
	inorder(root, &s)

	header := &TreeNode{}
	t := header
	for _, v := range s {
		t.Right = &TreeNode{Val: v}
		t = t.Right
	}
	return header.Right
}

func inorder(root *TreeNode, s *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, s)
	*s = append(*s, root.Val)
	inorder(root.Right, s)
}
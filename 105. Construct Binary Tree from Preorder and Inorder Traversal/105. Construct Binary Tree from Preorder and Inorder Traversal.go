/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <= 0 || len(inorder) <= 0 {
		return nil
	}

	v, preorder := preorder[0], preorder[1:]
	i := indexof(inorder, v)
	root := &TreeNode{Val: v}
	root.Left = buildTree(preorder[:i], inorder[:i])
	root.Right = buildTree(preorder[i:], inorder[i+1:])
	return root
}

func indexof(arr []int, s int) int {
	for i, v := range arr {
		if v == s {
			return i
		}
	}
	return -1
}
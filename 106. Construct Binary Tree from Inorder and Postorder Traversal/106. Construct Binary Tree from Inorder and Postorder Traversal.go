/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) <= 0 || len(postorder) <= 0 {
		return nil
	}

	j := len(postorder) - 1
	v := postorder[j]
	i := indexof(inorder, v)
	root := &TreeNode{Val: v}
	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:j])
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
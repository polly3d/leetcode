/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	m := 0
	maxDepth(root, &m)
	return m
}

func maxDepth(root *TreeNode, m *int) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left, m)
	r := maxDepth(root.Right, m)
	*m = max(*m, l+r)
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
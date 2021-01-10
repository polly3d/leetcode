/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := 1 + maxDepth(root.Left)
	r := 1 + maxDepth(root.Right)
	if l < r {
		return r
	}
	return l
}
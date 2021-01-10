/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	d := l - r
	if -1 <= d && d <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}

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
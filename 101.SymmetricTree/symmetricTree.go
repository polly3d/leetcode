package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return symmetric(root, root)
}

func symmetric(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l != nil && r != nil {
		if l.Val == r.Val {
			return symmetric(l.Left, r.Right) && symmetric(l.Right, r.Left)
		}
	}

	return false
}

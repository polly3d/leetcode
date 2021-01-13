/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return false
	}
	if traverse(s, t) {
		return true
	}
	if isSubtree(s.Left, t) || isSubtree(s.Right, t) {
		return true
	}
	return false
}

//Same tree
func traverse(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}

	if s.Val != t.Val {
		return false
	}

	return traverse(s.Left, t.Left) && traverse(s.Right, t.Right)
}
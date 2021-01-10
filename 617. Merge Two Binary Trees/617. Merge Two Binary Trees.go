package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	t := &TreeNode{}
	if t1 == nil && t2 != nil {
		t = t2
	}
	if t2 == nil && t1 != nil {
		t = t1
	}
	if t1 != nil && t2 != nil {
		t.Val = t1.Val + t2.Val
		t.Left = mergeTrees(t1.Left, t2.Left)
		t.Right = mergeTrees(t1.Right, t2.Right)
	}

	return t
}

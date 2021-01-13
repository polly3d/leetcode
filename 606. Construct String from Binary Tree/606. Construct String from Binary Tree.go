/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	s := fmt.Sprintf("%v", t.Val)
	if t.Left == nil && t.Right == nil {
		return s
	}
	s += fmt.Sprintf("(%v)", tree2str(t.Left))
	if t.Right != nil {
		s += fmt.Sprintf("(%v)", tree2str(t.Right))
	}
	return s
}
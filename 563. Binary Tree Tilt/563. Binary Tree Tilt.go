/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	r := 0
	postOrder(root, &r)
	return r
}

func postOrder(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	l := postOrder(root.Left, res)
	r := postOrder(root.Right, res)

	*res = *res + abs(l-r)
	return l + r + root.Val
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
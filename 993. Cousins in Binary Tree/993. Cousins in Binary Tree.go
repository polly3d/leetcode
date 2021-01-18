/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	_, dx, px := findNode(root, -1, x, 0)
	_, dy, py := findNode(root, -1, y, 0)
	return dx == dy && px != py
}

func findNode(root *TreeNode, parent, target, depth int) (bool, int, int) {
	if root == nil {
		return false, 0, 0
	}
	if root.Val == target {
		return true, depth, parent
	}
	if ex, d, p := findNode(root.Left, root.Val, target, depth+1); ex {
		return ex, d, p
	}
	if ex, d, p := findNode(root.Right, root.Val, target, depth+1); ex {
		return ex, d, p
	}
	return false, 0, 0
}
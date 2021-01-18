/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	s := []int{root.Val, math.MaxInt64}
	traverse(root, &s)
	if s[1] != math.MaxInt64 {
		return s[1]
	}
	return -1
}

func traverse(root *TreeNode, s *[]int) {
	if root == nil {
		return
	}
	t := *s
	if root.Val != t[0] {
		t[1] = min(root.Val, t[1])
	}
	traverse(root.Left, s)
	traverse(root.Right, s)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
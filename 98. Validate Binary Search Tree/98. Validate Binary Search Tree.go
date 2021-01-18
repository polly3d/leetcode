/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	s := []int{}
	inorder(root, &s)
	for i := len(s) - 1; i >= 1; i-- {
		if s[i] <= s[i-1] {
			return false
		}
	}
	return true
}

func inorder(root *TreeNode, s *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, s)
	*s = append(*s, root.Val)
	inorder(root.Right, s)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return traverse(root, math.MinInt64, math.MaxInt64)
}

func traverse(root *TreeNode, mi, ma int) bool {
	if root == nil {
		return true
	}
	if root.Val >= ma || root.Val <= mi {
		return false
	}

	return traverse(root.Left, mi, root.Val) && traverse(root.Right, root.Val, ma)
}
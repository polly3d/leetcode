/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return traverse(root, 0)
}

func traverse(root *TreeNode, currentSum int) int {
	currentSum = (currentSum << 1) | root.Val
	if root.Left == nil && root.Right == nil {
		return currentSum
	}

	total := 0
	if root.Left != nil {
		total += traverse(root.Left, currentSum)
	}
	if root.Right != nil {
		total += traverse(root.Right, currentSum)
	}

	return total
}

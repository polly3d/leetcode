package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	sum := 0
	inOrder(root, low, high, &sum)
	return sum
}

func inOrder(root *TreeNode, low, hight int, sum *int) {
	if root == nil {
		return
	}
	inOrder(root.Left, low, hight, sum)
	if low <= root.Val && root.Val <= hight {
		*sum = *sum + root.Val
	}
	inOrder(root.Right, low, hight, sum)
}

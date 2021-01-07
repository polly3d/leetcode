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
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

// func inorderTraversal(root *TreeNode) []int {
//     res := []int{}
//     inorder(root, &res)
//     return res
// }

// func inorder(root *TreeNode, output *[]int) {
//     if root == nil {
//         return
//     }
//     inorder(root.Left, output)
//     *output = append(*output, root.Val)
//     inorder(root.Right, output)
// }

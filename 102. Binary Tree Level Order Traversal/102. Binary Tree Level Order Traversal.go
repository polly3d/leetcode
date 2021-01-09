package leecode

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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{}
	queue = append(queue, root)

	res := [][]int{}
	for len(queue) > 0 {
		r := []int{}
		ln := queue[:len(queue)]
		queue = []*TreeNode{}
		for _, n := range ln {
			r = append(r, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		res = append(res, r)
	}
	return res
}

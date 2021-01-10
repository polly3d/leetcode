package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	s := []*TreeNode{root}
	for len(s) > 0 {
		ls := s[:]
		s = []*TreeNode{}
		r := []int{}
		for _, n := range ls {
			r = append(r, n.Val)
			if n.Left != nil {
				s = append(s, n.Left)
			}
			if n.Right != nil {
				s = append(s, n.Right)
			}
		}
		res = append([][]int{r}, res...)
	}
	return res
}

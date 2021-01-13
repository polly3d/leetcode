/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	m := map[int]int{}
	traverse(root, &m)
	max := -1
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	res := []int{}
	for k, v := range m {
		if v == max {
			res = append(res, k)
		}
	}
	return res
}

func traverse(root *TreeNode, m *map[int]int) {
	if root == nil {
		return
	}
	mm := *m
	mm[root.Val]++
	traverse(root.Left, m)
	traverse(root.Right, m)
}
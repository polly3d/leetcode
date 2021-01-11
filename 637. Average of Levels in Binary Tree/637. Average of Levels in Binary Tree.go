/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	res := []float64{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		sum := 0
		l := len(q)
		qt := q
		q = []*TreeNode{}
		for _, t := range qt {
			sum += t.Val
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		ave := float64(sum) / float64(l)
		res = append(res, ave)
	}
	return res
}
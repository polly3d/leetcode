/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return doGenerate(1, n)
}

func doGenerate(s, e int) []*TreeNode {
	if e < s {
		return []*TreeNode{nil}
	}
	res := []*TreeNode{}
	for i := s; i <= e; i++ {
		left := doGenerate(s, i-1)
		right := doGenerate(i+1, e)
		for _, l := range left {
			for _, r := range right {
				res = append(res, &TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	return res
}
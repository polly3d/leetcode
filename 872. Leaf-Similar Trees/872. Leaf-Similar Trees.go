/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	s1 := []int{}
	s2 := []int{}
	getLeaf(root1, &s1)
	getLeaf(root2, &s2)
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func getLeaf(root *TreeNode, s *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*s = append(*s, root.Val)
	}
	getLeaf(root.Left, s)
	getLeaf(root.Right, s)
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var l []int
	helper(root, &l)
	ret := math.MaxInt32
	for i := 1; i < len(l); i++ {
		diff := l[i] - l[i-1]
		if diff < ret {
			ret = diff
		}
	}
	return ret
}

func helper(root *TreeNode, l *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, l)
	*l = append(*l, root.Val)
	helper(root.Right, l)
}

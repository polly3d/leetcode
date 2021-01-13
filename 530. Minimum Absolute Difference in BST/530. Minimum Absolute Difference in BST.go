/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	lastNode := -1
	min := math.MaxInt32
	inorder(root, &lastNode, &min)
	return min
}

func inorder(root *TreeNode, lastNode *int, min *int) {
	if root != nil {
		inorder(root.Left, lastNode, min)
		if *lastNode != -1 {
			*min = mini(*min, abs(root.Val-*lastNode))
		}
		*lastNode = root.Val
		inorder(root.Right, lastNode, min)
	}
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func mini(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for (root.Val-p.Val)*(root.Val-q.Val) > 0 {
		if root.Val > p.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	mi := min(p.Val, q.Val)
	ma := max(p.Val, q.Val)
	if mi <= root.Val && root.Val <= ma {
		return root
	}
	if root.Val > mi {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return lowestCommonAncestor(root.Right, p, q)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
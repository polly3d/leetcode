/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	depth := 0
	for _, n := range root.Children {
		depth = max(depth, maxDepth(n))
	}
	return 1 + depth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
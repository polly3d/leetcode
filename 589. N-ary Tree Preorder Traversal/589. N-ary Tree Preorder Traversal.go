/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	s := []int{root.Val}
	preorderChildren(root.Children, &s)
	return s
}

func preorderChildren(ns []*Node, s *[]int) {
	if ns == nil {
		return
	}
	for _, c := range ns {
		*s = append(*s, c.Val)
		preorderChildren(c.Children, s)
	}
}
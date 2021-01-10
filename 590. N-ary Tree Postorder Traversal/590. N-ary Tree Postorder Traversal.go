/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	s := []int{}
	postorderChildren(root.Children, &s)
	s = append(s, root.Val)
	return s
}

func postorderChildren(ns []*Node, s *[]int) {
	if ns == nil {
		return
	}
	for _, c := range ns {
		postorderChildren(c.Children, s)
		*s = append(*s, c.Val)
	}
}
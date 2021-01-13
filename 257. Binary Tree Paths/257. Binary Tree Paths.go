/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	lines := []string{}
	dfs(root, &lines, []string{})
	return lines
}

func dfs(root *TreeNode, lines *[]string, line []string) {
	if root == nil {
		return
	}
	line = append(line, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		//Add node
		*lines = append(*lines, strings.Join(line, "->"))
	}
	dfs(root.Left, lines, line)
	dfs(root.Right, lines, line)
}
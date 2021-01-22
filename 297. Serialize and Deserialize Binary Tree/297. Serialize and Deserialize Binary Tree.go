/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

const split = ","
const ns = "x"

type Codec struct {
	data []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	s := ""
	preorder(root, &s)
	return s
}

func preorder(root *TreeNode, s *string) {
	if root == nil {
		*s = *s + ns + split
		return
	}
	*s = *s + strconv.Itoa(root.Val) + split
	preorder(root.Left, s)
	preorder(root.Right, s)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	ss := strings.Split(data, split)
	this.data = ss
	return this.buildTree()
}

func (this *Codec) buildTree() *TreeNode {
	c := this.data[0]
	this.data = this.data[1:]
	if c == ns {
		return nil
	}
	v, _ := strconv.Atoi(c)
	root := &TreeNode{Val: v}
	root.Left = this.buildTree()
	root.Right = this.buildTree()
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
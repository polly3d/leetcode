package structs

type ListNode struct {
	Val  int
	Next *ListNode
}

func IntsToList(nums []int) *ListNode {
	head := &ListNode{}
	dummy := head
	for _, v := range nums {
		node := &ListNode{Val: v}
		dummy.Next = node
		dummy = node
	}
	return head.Next
}

func ListToInt(node *ListNode) []int {
	res := []int{}
	for {
		if node == nil {
			break
		}
		res = append(res, node.Val)
		node = node.Next
	}
	return res
}

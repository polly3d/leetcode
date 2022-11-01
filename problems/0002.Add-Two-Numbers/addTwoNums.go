package leetcode

import (
	"leetcode/structs"
)

type ListNode = structs.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: -1}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry > 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		carry = (n1 + n2 + carry) / 10
		current = current.Next
	}
	return head.Next
}

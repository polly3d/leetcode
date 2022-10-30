package leetcode

import (
	"leetcode/structs"
)

func addTwoNumbers(l1 *structs.ListNode, l2 *structs.ListNode) *structs.ListNode {
	dummy := &structs.ListNode{Val: -1}
	pre := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil && l2 != nil {
			v := carry + l1.Val + l2.Val
			sum = v % 10
			carry = v / 10
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			v := carry + l1.Val
			sum = v % 10
			carry = v / 10
			l1 = l1.Next

		} else if l2 != nil {
			v := carry + l2.Val
			sum = v % 10
			carry = v / 10
			l2 = l2.Next
		}
		pre.Next = &structs.ListNode{
			Val: sum,
		}
		pre = pre.Next
	}
	if carry > 0 {
		pre.Next = &structs.ListNode{
			Val: carry,
		}
	}
	return dummy.Next
}

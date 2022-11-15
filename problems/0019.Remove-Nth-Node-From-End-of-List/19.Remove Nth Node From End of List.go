package leetcode

import "leetcode/structs"

func removeNthFromEnd(head *structs.ListNode, n int) *structs.ListNode {
	dummy := &structs.ListNode{
		Next: head,
	}
	fast, slow := dummy, dummy
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

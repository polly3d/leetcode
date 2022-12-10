package leetcode

import "leetcode/structs"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *structs.ListNode, k int) *structs.ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

// 前插法
// 1->2->3->4
// 2->3->4, 1->4
// 3->4, 2->1->4
// 3->2->1->4
func reverse(first *structs.ListNode, last *structs.ListNode) *structs.ListNode {
	prev := last
	for first != last {
		tmp := first.Next
		first.Next = prev
		prev = first
		first = tmp
	}
	return prev
}

// 尾插法
func reverseKGroup2(head *structs.ListNode, k int) *structs.ListNode {
	dummy := &structs.ListNode{Next: head}
	pre := dummy
	tail := dummy
	for {
		count := k
		for count > 0 && tail != nil {
			count--
			tail = tail.Next
		}
		if tail == nil {
			break
		}
		head := pre.Next
		for pre.Next != tail {
			cur := pre.Next
			pre.Next = cur.Next
			cur.Next = tail.Next
			tail.Next = cur
		}
		pre = head
		tail = head
	}
	return dummy.Next
}

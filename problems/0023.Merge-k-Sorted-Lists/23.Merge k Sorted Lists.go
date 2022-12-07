package leetcode

import "leetcode/structs"

func mergeKLists(lists []*structs.ListNode) *structs.ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	res := &structs.ListNode{}
	for _, l := range lists {
		res = mergeTwoLists(res, l)
	}
	return res.Next
}

func mergeTwoLists(list1 *structs.ListNode, list2 *structs.ListNode) *structs.ListNode {
	dummy := &structs.ListNode{}
	head := dummy
	for list1 != nil || list2 != nil {
		n := &structs.ListNode{}
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				n.Val = list1.Val
				list1 = list1.Next
			} else {
				n.Val = list2.Val
				list2 = list2.Next
			}
		} else if list1 == nil {
			n.Val = list2.Val
			list2 = list2.Next
		} else if list2 == nil {
			n.Val = list1.Val
			list1 = list1.Next
		}
		head.Next = n
		head = head.Next
	}
	return dummy.Next
}

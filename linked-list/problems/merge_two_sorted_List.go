package linkedlist

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = list1
			curr = curr.Next

			list1 = list1.Next
		} else {
			curr.Next = list2
			curr = curr.Next

			list2 = list2.Next
		}
	}

	if list1 != nil {
		curr.Next = list1
	} else if list2 != nil {
		curr.Next = list2
	}

	return dummy.Next
}

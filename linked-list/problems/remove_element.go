package linkedlist

func RemoveElements(head *ListNode, val int) *ListNode {
	prev := &ListNode{Next: head}
	current := prev
	for current != nil && current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return prev.Next
}

package linkedlist

func DeleteNode(node *ListNode) {
	curr := node

	curr.Val = curr.Next.Val
	curr.Next = curr.Next.Next
}

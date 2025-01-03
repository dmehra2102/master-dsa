package linkedlist

func IsPalindromeLL(head *ListNode) bool {
	// STEP-1 : With the help of slow and fast pointer find out mid point
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// STEP-2 : Reverse the 2nd half linked list
	var prev *ListNode
	curr := slow
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// STEP-3 : Now do the comparision using 2 pointers again
	left := head
	right := prev

	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}

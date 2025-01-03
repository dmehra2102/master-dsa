package linkedlist

// SOLUTION-1
func GetIntersectionNode1(headA, headB *ListNode) *ListNode {
	// We can also use hash table for solving this problem but it wi;; take SC => O(n)

	// STEP-1 : Find the length of Linked Lists
	curr1 := headA
	curr2 := headB

	len1 := 0
	len2 := 0

	for curr1 != nil || curr2 != nil {
		if curr1 != nil {
			len1++
			curr1 = curr1.Next
		}

		if curr2 != nil {
			len2++
			curr2 = curr2.Next
		}
	}

	// STEP-2 : Move one pointer to the difference of length
	head1 := headA
	head2 := headB

	var diff int

	if len1 < len2 {
		diff = len2 - len1
		for diff != 0 {
			head2 = head2.Next
			diff--
		}
	} else {
		diff = len1 - len2
		for diff != 0 {
			head1 = head1.Next
			diff--
		}
	}

	// STEP-3 with 2 pointers find out intersection
	for head1 != nil {
		if head1 == head2 {
			return head1
		}
		head1 = head1.Next
		head2 = head2.Next
	}

	return nil

}

// ------------------------------------------------------

// SOLUTION -2
func GetIntersectionNode2(headA, headB *ListNode) *ListNode {
	curr1, curr2 := headA, headB

	// Traverse both lists to find their lengths
	for curr1 != curr2 {
		curr1 = next(curr1, headB)
		curr2 = next(curr2, headA)
	}

	return curr1
}

// Helper function to move to the next node or switch to the other list
func next(curr *ListNode, otherHead *ListNode) *ListNode {
	if curr == nil {
		return otherHead
	}
	return curr.Next
}

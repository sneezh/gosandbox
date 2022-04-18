package linked_list

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	if headA == headB {
		return headA
	}

	originalBHead := headB

	for headA != nil {
		for headB != nil {
			if headB == headA {
				return headB
			}

			headB = headB.Next
		}
		headB = originalBHead
		headA = headA.Next
	}

	return nil
}

func getIntersectionNodeFaster(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a := headA
	b := headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

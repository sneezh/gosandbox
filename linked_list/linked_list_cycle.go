package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycleTwoPointers(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	p1, p2 := head, head.Next
	for p1 != p2 {
		if p2 == nil || p2.Next == nil {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return true
}

func HasCycleHash(head *ListNode) bool {
	existingNodes := make(map[*ListNode]struct{})
	for node := head; node != nil; node = node.Next {
		if _, exists := existingNodes[node]; exists {
			return true
		}
		existingNodes[node] = struct{}{}
	}

	return false
}

func DetectCycle(head *ListNode) *ListNode {
	sp, fp := head, head
	for fp != nil && fp.Next != nil {
		sp = sp.Next
		fp = fp.Next.Next
		if sp == fp {
			break
		}
	}
	if fp == nil || fp.Next == nil {
		return nil
	}
	fp = head

	for fp != sp {
		fp = fp.Next
		sp = sp.Next
	}
	return fp
}

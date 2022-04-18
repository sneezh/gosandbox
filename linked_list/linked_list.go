package linked_list

type MyLinkedList struct {
	Head *Node
}

type Node struct {
	Val  int
	Next *Node
}

func (ll *MyLinkedList) Get(index int) int {
	var count int

	for node := ll.Head; node != nil; node = node.Next {
		if count == index {
			return node.Val
		}
		count++
	}
	return -1
}

func (ll *MyLinkedList) AddAtHead(val int) {
	ll.Head = &Node{val, ll.Head}
}

func (ll *MyLinkedList) AddAtTail(val int) {
	if ll.Head == nil {
		ll.Head = &Node{Val: val}
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = &Node{Val: val}
}

func (ll *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		ll.AddAtHead(val)
		return
	}

	curr := ll.Head
	for i := 1; i < index; i++ {
		curr = curr.Next
		if curr == nil {
			return
		}
	}

	if curr == nil {
		return
	}
	curr.Next = &Node{Val: val, Next: curr.Next}
}

func (ll *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		if ll.Head == nil {
			return
		}

		ll.Head = ll.Head.Next
		return
	}

	var count int
	prev := ll.Head
	for node := ll.Head; node != nil; node = node.Next {
		if count == index {
			prev.Next = node.Next
			break
		}
		prev = node
		count++
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy

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

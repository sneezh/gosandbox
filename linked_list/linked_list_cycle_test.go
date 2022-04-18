package linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasCycleTwoPointers(t *testing.T) {
	listNode := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  -4,
					Next: nil,
				},
			},
		},
	}

	listNode.Next.Next.Next.Next = listNode.Next

	t.Run("true", func(t *testing.T) {
		assert.Equal(t, true, HasCycleTwoPointers(nil))
	})

	listNode = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  -4,
					Next: nil,
				},
			},
		},
	}

	listNode.Next.Next.Next.Next = listNode.Next

	t.Run("true", func(t *testing.T) {
		assert.Equal(t, false, HasCycleTwoPointers(nil))
	})
}

func TestHasCycleHash(t *testing.T) {
	listNode := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  -4,
					Next: nil,
				},
			},
		},
	}

	listNode.Next.Next.Next.Next = listNode.Next

	t.Run("true", func(t *testing.T) {
		assert.Equal(t, true, HasCycleHash(listNode))
	})

	listNode = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  -4,
					Next: nil,
				},
			},
		},
	}

	t.Run("false", func(t *testing.T) {
		assert.Equal(t, false, HasCycleHash(listNode))
	})
}

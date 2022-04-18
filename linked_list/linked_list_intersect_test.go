package linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListIntersect(t *testing.T) {
	listNodeIntersectPart := &ListNode{
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

	listNodeA := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  1,
			Next: listNodeIntersectPart,
		},
	}

	listNodeB := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val:  6,
				Next: listNodeIntersectPart,
			},
		},
	}

	t.Run("true", func(t *testing.T) {
		assert.Equal(t, listNodeIntersectPart, getIntersectionNode(listNodeA, listNodeB))
	})
}

func TestListIntersectFaster(t *testing.T) {
	listNodeIntersectPart := &ListNode{
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

	listNodeA := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val:  2,
						Next: listNodeIntersectPart,
					},
				},
			},
		},
	}

	listNodeB := &ListNode{
		Val:  6,
		Next: listNodeIntersectPart,
	}

	t.Run("true", func(t *testing.T) {
		assert.Equal(t, listNodeIntersectPart, getIntersectionNodeFaster(listNodeA, listNodeB))
	})
}

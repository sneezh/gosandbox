package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSemmetric(t *testing.T) {
	var tests = []struct {
		name     string
		expected bool
		given    *TreeNode
	}{
		{
			"symmetric",
			true,
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
		},
		{
			"diffrent",
			false,
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, isSymmetric(tt.given))
		})
	}
}

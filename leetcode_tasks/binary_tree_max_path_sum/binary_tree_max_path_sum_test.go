package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxPathSum(t *testing.T) {
	rootNode := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}

	assert.Equal(t, 6, maxPathSum(rootNode))
}

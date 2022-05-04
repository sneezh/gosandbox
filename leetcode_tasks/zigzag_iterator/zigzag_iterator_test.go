package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZigzagIterator(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, runZigzag([]int{1, 3, 5}, []int{2, 4, 6}))
	assert.Equal(t, []int{1, 2, 4, 6}, runZigzag([]int{1}, []int{2, 4, 6}))
	assert.Equal(t, []int{}, runZigzag([]int{}, []int{}))
	assert.Equal(t, []int{1, 3, 2}, runZigzag([]int{1, 2}, []int{3}))
}

func TestZigzagWithSwap(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, runZigzagWithSwap([]int{1, 3, 5}, []int{2, 4, 6}))
	assert.Equal(t, []int{1, 2, 4, 6}, runZigzagWithSwap([]int{1}, []int{2, 4, 6}))
	assert.Equal(t, []int{}, runZigzagWithSwap([]int{}, []int{}))
	assert.Equal(t, []int{1, 3, 2}, runZigzagWithSwap([]int{1, 2}, []int{3}))
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{4, 5, 6}
	merge(nums1, 3, nums2, 3)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, nums1)
}

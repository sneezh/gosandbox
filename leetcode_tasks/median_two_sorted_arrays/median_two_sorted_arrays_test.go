package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	var tests = []struct {
		name     string
		expected float64
		nums1    []int
		nums2    []int
	}{
		{"even count", 3.5, []int{1, 2, 3}, []int{4, 5, 6}},
		{"odd count", 4, []int{1, 2, 3}, []int{4, 5, 6, 7}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, findMedianSortedArrays(tt.nums1, tt.nums2))
		})
	}
}

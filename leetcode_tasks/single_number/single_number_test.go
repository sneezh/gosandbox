package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinleNumber(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		nums     []int
	}{
		{
			name:     "one pair",
			expected: 1,
			nums:     []int{1, 2, 2},
		},
		{
			name:     "last element",
			expected: 4,
			nums:     []int{1, 1, 2, 2, 4},
		},
		{
			name:     "one value",
			expected: 1,
			nums:     []int{1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, singleNumber(test.nums))
		})
	}
}

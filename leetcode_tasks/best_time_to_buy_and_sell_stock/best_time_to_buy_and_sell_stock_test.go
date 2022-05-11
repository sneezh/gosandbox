package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		prices   []int
	}{
		{
			name:     "default",
			expected: 5,
			prices:   []int{7, 1, 5, 3, 6, 4},
		},
		{
			name:     "no profit",
			expected: 0,
			prices:   []int{6, 4, 3, 2, 1},
		},
		{
			name:     "empty prices",
			expected: 0,
			prices:   []int{},
		},
		{
			name:     "same prices",
			expected: 0,
			prices:   []int{3, 3, 3, 3, 3, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, maxProfit(test.prices))
		})
	}
}

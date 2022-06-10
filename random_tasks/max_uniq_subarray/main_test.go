package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMaxLen(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		given    []int
	}{
		{
			"empty",
			0,
			[]int{},
		},
		{
			"one elem",
			1,
			[]int{1},
		},
		{
			"all uniq",
			4,
			[]int{1, 2, 3, 4},
		},
		{
			"simple",
			4,
			[]int{1, 2, 3, 4, 2},
		},
		{
			"repeated elem",
			5,
			[]int{1, 2, 3, 4, 1, 5},
		},
		{
			"",
			5,
			[]int{1, 2, 3, 4, 5, 5, 1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, getMaxLen(test.given))
		})
	}
}

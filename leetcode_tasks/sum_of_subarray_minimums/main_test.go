package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumSubarrayMins(t *testing.T) {
	var tests = []struct {
		name     string
		expected int
		given    []int
	}{
		{"t1", 17, []int{3, 1, 2, 4}},
		//		{"t2", 444, []int{11, 81, 94, 43, 3}},
		//		{"empty", 0, []int{}},
		//		{"one elem", 5, []int{5}},
		//		{"same values", 10, []int{1, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, sumSubarrayMins(tt.given))
		})
	}
}

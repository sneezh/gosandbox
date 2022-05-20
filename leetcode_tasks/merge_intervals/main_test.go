package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	var tests = []struct {
		name     string
		given    [][]int
		expected [][]int
	}{
		{"many intervals", [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{"not sorted", [][]int{{2, 6}, {1, 3}, {15, 18}, {8, 10}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{"can be one", [][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{"empty", [][]int{}, [][]int{}},
		{"cannot merge", [][]int{{1, 4}, {5, 10}}, [][]int{{1, 4}, {5, 10}}},
		{"included interval", [][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, merge(tt.given))
		})
	}
}

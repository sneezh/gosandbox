package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinMeetingRooms(t *testing.T) {
	tests := []struct {
		name      string
		expected  int
		intervals [][]int
	}{
		{
			name:      "default",
			expected:  2,
			intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
		},
		{
			name:      "only one",
			expected:  1,
			intervals: [][]int{{7, 10}, {2, 4}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, minMeetingRooms(test.intervals))
		})
	}
}

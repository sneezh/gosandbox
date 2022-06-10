package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		name     string
		given    string
		expected [][]string
	}{
		{
			name:     "default",
			given:    "aab",
			expected: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			name:  "with append",
			given: "cbbbcc",
			expected: [][]string{
				{"c", "b", "b", "b", "c", "c"},
				{"c", "b", "b", "b", "cc"},
				{"c", "b", "bb", "c", "c"},
				{"c", "b", "bb", "cc"}, {"c", "bb", "b", "c", "c"},
				{"c", "bb", "b", "cc"},
				{"c", "bbb", "c", "c"},
				{"c", "bbb", "cc"},
				{"cbbbc", "c"},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, partition(test.given))
		})
	}
}

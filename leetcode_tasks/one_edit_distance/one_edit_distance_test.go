package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsOneEditDistance(t *testing.T) {
	tests := []struct {
		name     string
		s, t     string
		expected bool
	}{
		{
			name:     "1",
			s:        "ab",
			t:        "acb",
			expected: true,
		},
		{
			name:     "2",
			s:        "ab",
			t:        "acba",
			expected: false,
		},
		{
			name:     "3",
			s:        "",
			t:        "",
			expected: false,
		},
		{
			name:     "3",
			s:        "aa",
			t:        "aa",
			expected: false,
		},
		{
			name:     "4",
			s:        "a",
			t:        "A",
			expected: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, isOneEditDistance(test.s, test.t))
		})
	}
}

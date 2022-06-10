package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTopKFrequentWords(t *testing.T) {
	tests := []struct {
		name     string
		expected []string
		words    []string
		k        int
	}{
		{
			name:     "empty",
			expected: []string{},
			words:    []string{},
			k:        0,
		},
		{
			name:     "aaa",
			expected: []string{"a"},
			words:    []string{"a", "aa", "aaa"},
			k:        1,
		},
		{
			name:     "ilovecode",
			expected: []string{"love", "code"},
			words:    []string{"i", "love", "code", "love"},
			k:        2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, topKFrequent(test.words, test.k))
		})
	}
}

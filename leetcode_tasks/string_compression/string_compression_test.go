package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompress(t *testing.T) {
	tests := []struct {
		name     string
		chars    []byte
		expected int
	}{
		{
			name:     "default",
			chars:    []byte("aabbccc"),
			expected: 6,
		},
		{
			name:     "default",
			chars:    []byte("abbbbbbbbbbbb"),
			expected: 4,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, compress(test.chars))
		})
	}
}

func TestCompressResult(t *testing.T) {
	tests := []struct {
		name     string
		chars    []byte
		expected string
	}{
		{
			name:     "default",
			chars:    []byte("aabbccc"),
			expected: "a2b2c3",
		},
		{
			name:     "default",
			chars:    []byte("abbbbbbbbbbbb"),
			expected: "ab12",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, compressResult(test.chars))
		})
	}
}

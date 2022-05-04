package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddStrings(t *testing.T) {
	tests := []struct {
		name                 string
		num1, num2, expected string
	}{
		{
			name:     "1",
			num1:     "1",
			num2:     "2",
			expected: "3",
		},
		{
			name:     "2",
			num1:     "1999",
			num2:     "2",
			expected: "2001",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, addStrings(test.num1, test.num2))
		})
	}
}

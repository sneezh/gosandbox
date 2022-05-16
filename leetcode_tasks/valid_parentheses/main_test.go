package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	var tests = []struct {
		name     string
		expected bool
		given    string
	}{
		//	{"default", true, "()"},
		//	{"three types", true, "(){}[]"},
		//	{"invalid", false, "(]"},
		{"many opens all closes", true, "{{{[[[((()))]]]}}}"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, isValid(tt.given))
		})
	}
}

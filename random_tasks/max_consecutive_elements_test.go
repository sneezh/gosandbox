package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxConsecutiveElements(t *testing.T) {
	assert.Equal(t, 1, maxConsecutiveElements("avbcs"))
	assert.Equal(t, 5, maxConsecutiveElements("aaaaa"))
}

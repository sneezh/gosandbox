package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, 4, seach([]int{-1, 0, 3, 5, 9, 11}, 9))
}

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubarraySum(t *testing.T) {
	assert.Equal(t, true, subarraySum([]int{1, 2, 3, 4}, 4))
	assert.Equal(t, true, subarraySum([]int{1, 2, 3, 4}, 7))
	assert.Equal(t, true, subarraySum([]int{1, 2, 3, 4}, 10))
	assert.Equal(t, false, subarraySum([]int{1, 2, 3, 4}, 11))
	assert.Equal(t, false, subarraySum([]int{1, 2, 3, 4}, 8))
}

package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		assert.Equal(t, []int{0, 1, 9, 16, 100}, sortedSquares([]int{-4, -1, 0, 3, 10}))
	})
}

func TestAppend(t *testing.T) {
	clearSlice := []int{1, 2, 3}
	clearSlice = append(clearSlice[0:0], clearSlice[1:]...)
	assert.Equal(t, []int{2, 3}, clearSlice)
}

func TestDeleting(t *testing.T) {
	clearSlice := []int{1}
	clearSlice = clearSlice[1:]
	assert.Equal(t, []int{}, clearSlice)
}

func TestSwap(t *testing.T) {
	swap := func(arr []int) []int {
		arr = append(arr[1:], arr[0])
		return arr
	}
	arr := []int{1, 2, 3, 4, 5}
	arr = swap(arr)
	assert.Equal(t, []int{2, 3, 4, 5, 1}, arr)

	arr = append(arr[1:], arr[0])
	assert.Equal(t, []int{3, 4, 5, 1, 2}, arr)
}

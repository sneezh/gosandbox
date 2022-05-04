package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	a := minMeetingRooms([][]int{{7, 10}, {2, 4}})
	fmt.Println(a)
	return
}

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1]
	})

	h := &MinHeap{}
	for _, interval := range intervals {
		if h.Len() > 0 && h.Top().(int) <= interval[0] {
			heap.Pop(h)
		}
		heap.Push(h, interval[1])
	}

	return h.Len()
}

type MinHeap []int

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Top() interface{} {
	return (*h)[0]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	n := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return n
}

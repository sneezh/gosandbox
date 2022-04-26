package main

import (
	"container/heap"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) <= 1 {
		return len(intervals)
	}

	ans := 0
	var end []int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})

	for _, v := range intervals {
		if len(end) == 0 {
			end = append(end, v[1])
			continue
		}

		if v[0] >= end[0] {
			end = end[1:]
		}
		end = append(end, v[1])
		sort.Ints(end)
		if len(end) > ans {
			ans = len(end)
		}

	}
	return ans
}

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

// Min heap.
func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	n := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return n
}

func (h *Heap) Top() interface{} {
	return (*h)[0]
}

func minMeetingRoomsWithMinHeap(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})

	h := &Heap{}

	for _, interval := range intervals {

		if h.Len() > 0 && h.Top().(int) <= interval[0] {
			heap.Pop(h)
		}
		heap.Push(h, interval[1])
	}

	return h.Len()
}

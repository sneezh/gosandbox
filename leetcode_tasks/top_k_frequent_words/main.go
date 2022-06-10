package main

import "container/heap"

type wordCount struct {
	word  string
	count int
}

type minHeap []wordCount

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	if h[i].count == h[j].count {
		return h[i].word > h[j].word
	}

	return h[i].count < h[j].count
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h minHeap) Top() wordCount {
	return h[0]
}

func (h *minHeap) Pop() interface{} {
	n := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return n
}

func (h *minHeap) Push(v interface{}) {
	*h = append(*h, v.(wordCount))
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, v := range words {
		m[v]++
	}

	h := minHeap{}
	heap.Init(&h)
	for word, count := range m {
		if len(h) < k {
			heap.Push(&h, wordCount{
				word:  word,
				count: count,
			})

			continue
		}

		if count > h.Top().count || count == h.Top().count && word < h.Top().word {
			heap.Pop(&h)
			heap.Push(&h, wordCount{
				word:  word,
				count: count,
			})
		}
	}

	res := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(&h).(wordCount).word
	}

	return res
}

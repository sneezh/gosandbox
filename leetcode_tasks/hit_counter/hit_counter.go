package hit_counter

import "math"

type HitCounter struct {
	hits map[int]int
}

func Constructor() HitCounter {
	return HitCounter{
		hits: make(map[int]int, 300),
	}
}

func (hc *HitCounter) Hit(timestamp int) {
	hc.hits[timestamp]++
}

func (hc *HitCounter) GetHits(timestamp int) int {
	hits := 0
	start := int(math.Max(0, float64(timestamp-299)))

	for i := start; i <= timestamp; i++ {
		v, ok := hc.hits[i]
		if ok {
			hits += v
		}
	}

	return hits
}

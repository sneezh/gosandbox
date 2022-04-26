package main

type ZigzagIterator struct {
	arrays         [][]int
	currentPointer int
}

func Constructor(input ...[]int) *ZigzagIterator {
	var arrays [][]int
	for _, v := range input {
		if len(v) > 0 {
			arrays = append(arrays, v)
		}
	}
	return &ZigzagIterator{
		arrays:         arrays,
		currentPointer: 0,
	}
}

func (zi *ZigzagIterator) hasNext() bool {
	return len(zi.arrays) > 0
}

func (zi *ZigzagIterator) next() int {
	if !zi.hasNext() {
		return -1
	}

	next := zi.arrays[zi.currentPointer][0]

	zi.arrays[zi.currentPointer] = zi.arrays[zi.currentPointer][1:]

	if len(zi.arrays[zi.currentPointer]) == 0 {
		if len(zi.arrays) == 1 {
			zi.arrays = [][]int{}
		} else if zi.currentPointer == 0 {
			zi.arrays = zi.arrays[1:]
		} else if zi.currentPointer == len(zi.arrays)-1 {
			zi.arrays = zi.arrays[:zi.currentPointer]
		} else {
			zi.arrays = append(zi.arrays[0:zi.currentPointer-1], zi.arrays[zi.currentPointer:]...)
		}
	}

	zi.currentPointer++
	if zi.currentPointer >= len(zi.arrays) {
		zi.currentPointer = 0
	}

	return next
}

func (zi *ZigzagIterator) nextWithSwap() int {
	if !zi.hasNext() {
		return -1
	}

	next := zi.arrays[0][0]

	zi.arrays[0] = zi.arrays[0][1:]

	if len(zi.arrays[0]) > 0 {
		zi.arrays = append(zi.arrays[1:], zi.arrays[0])
	} else {
		zi.arrays = zi.arrays[1:]
	}

	return next
}

func runZigzag(v1, v2 []int) []int {
	ans := []int{}
	zi := Constructor(v1, v2)
	for zi.hasNext() {
		ans = append(ans, zi.next())
	}

	return ans
}

func runZigzagWithSwap(v1, v2 []int) []int {
	ans := []int{}
	zi := Constructor(v1, v2)
	for zi.hasNext() {
		ans = append(ans, zi.nextWithSwap())
	}

	return ans
}

package main

func singleNumberHashMap(nums []int) int {
	m := map[int]struct{}{}

	for _, v := range nums {
		if _, ok := m[v]; ok {
			delete(m, v)
		} else {
			m[v] = struct{}{}
		}
	}

	for t := range m {
		return t
	}

	return -1
}

func singleNumber(nums []int) int {
	a := 0
	for _, v := range nums {
		a ^= v
	}

	return a
}

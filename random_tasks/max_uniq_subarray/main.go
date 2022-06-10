package main

func getMaxLen(arr []int) int {
	m := map[int]int{}
	prev := 0
	max := 0
	for k, v := range arr {
		prev = maxInt(prev, m[v])
		max = maxInt(max, k-prev+1)
		m[v] = k + 1
	}

	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

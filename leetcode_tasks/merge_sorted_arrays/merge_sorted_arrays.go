package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1

	for p := m + n - 1; p >= 0; p-- {
		if p2 < 0 {
			break
		}

		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
	}
}

package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)

	if n+m == 0 {
		return 0.0
	}

	mergedArr := make([]int, 0, m+n)
	mergedArr = append(mergedArr, nums1...)
	mergedArr = append(mergedArr, make([]int, m)...)
	p1, p2 := n-1, m-1

	for p := m + n - 1; p >= 0; p-- {
		if p2 < 0 {
			break
		}

		if p1 >= 0 && mergedArr[p1] > nums2[p2] {
			mergedArr[p] = mergedArr[p1]
			p1--
		} else {
			mergedArr[p] = nums2[p2]
			p2--
		}
	}

	if (n+m)%2 != 0 {
		return float64(mergedArr[(n+m)/2])
	} else {
		return float64((float64(mergedArr[(n+m)/2-1] + mergedArr[(n+m)/2])) / 2)
	}
}

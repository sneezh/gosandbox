package main

import "fmt"

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}

	l, r := 0, 0
	var res int
	for r < len(chars) {
		r++

		if r == len(chars) || chars[l] != chars[r] {
			count := r - l
			res++

			if count > 1 {
				tmpCount := fmt.Sprintf("%d", count)
				res += len([]byte(tmpCount))
			}
			l = r
		}
	}

	return res
}

func compressResult(chars []byte) string {
	if len(chars) == 0 {
		return ""
	}

	l, r := 0, 0
	var res []byte
	for r < len(chars) {
		r++

		if r == len(chars) || chars[l] != chars[r] {
			count := r - l
			res = append(res, chars[l])

			if count > 1 {
				tmpCount := fmt.Sprintf("%d", count)
				res = append(res, []byte(tmpCount)...)
			}
			l = r
		}
	}

	return string(res)
}

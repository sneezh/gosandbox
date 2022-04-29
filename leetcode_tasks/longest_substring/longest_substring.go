package main

func lengthOfLongestSubstringTwoDistinct(s string) int {
	if len(s) < 3 {
		return len(s)
	}

	dict := make(map[byte]int)
	l, r := 0, 0
	max := 1
	dict[s[0]]++

	for r < len(s)-1 {
		r++
		dict[s[r]]++
		if len(dict) <= 2 {
			max = maxInt(max, r-l+1)
		}
		for len(dict) > 2 {
			dict[s[l]]--
			if dict[s[l]] == 0 {
				delete(dict, s[l])
			}
			l++
		}
	}

	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func lengthOfLongestSubstringTwoDistinctApproach2(s string) int {
	if len(s) < 3 {
		return len(s)
	}

	dict := make(map[byte]int)
	l, r := 0, 0
	max := 0

	for r < len(s) {
		dict[s[r]] = r
		if len(dict) == 3 {
			smallIdx := getSmallestValue(dict)
			delete(dict, s[smallIdx])
			l = smallIdx + 1
		}
		max = maxInt(max, r-l+1)
		r++
	}

	return max
}

func getSmallestValue(m map[byte]int) (smallest int) {
	for _, smallest = range m {
		break
	}

	for _, v := range m {
		if v < smallest {
			smallest = v
		}
	}

	return smallest
}

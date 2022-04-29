package main

func isOneEditDistance(s, t string) bool {
	if s == t {
		return false
	}

	sl := len(s)
	tl := len(t)

	if sl > tl {
		return isOneEditDistance(t, s)
	}

	if tl-sl > 1 {
		return false
	}

	for i := 0; i < sl; i++ {
		if s[i] != t[i] {
			if sl == tl {
				return s[i+1:] == t[i+1:]
			} else {
				return s[i:] == t[i+1:]
			}
		}
	}

	return sl+1 == tl
}

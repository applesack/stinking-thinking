package p3

func lengthOfLongestSubstring(s string) int {
	var cache = [127]int{}
	for i := range cache {
		cache[i] = -1
	}
	var max, distance = 0, 0
	for i, ch := range s {
		distance = i - cache[ch]
		if distance > i {
			cache[ch] = i
		} else {
			if distance > max {
				max = distance
			}
		}
	}
	return max
}

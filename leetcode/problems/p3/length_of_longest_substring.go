package p3

// 入参包含有大小写英文字母,下划线,符号和空格
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	var mem = [127]int{}
	for i := range mem {
		mem[i] = -1
	}
	max, l, r, distance := 0, 0, 0, 0
	for r < len(s) {
		idx := mem[s[r]]
		if idx >= l {
			l = idx + 1
		}
		distance = r - l + 1
		if distance > max {
			max = distance
		}
		mem[s[r]] = r
		r++
	}
	return max
}

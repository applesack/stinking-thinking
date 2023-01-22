package p409

func longestPalindrome(s string) int {
	m := make(map[int32]int)
	for _, ch := range s {
		m[ch] += 1
	}
	single, pair := 0, 0
	for _, cnt := range m {
		if cnt%2 == 1 {
			single = 1
		}
		pair += cnt >> 1 << 1
	}
	return single + pair
}

package p9

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	for l, r := 0, len(s)-1; l < len(s)/2; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}

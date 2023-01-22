package p415

import (
	"strings"
)

func addStrings(num1 string, num2 string) string {
	ans := strings.Builder{}
	len1, len2 := len(num1), len(num2)
	carry, sum := uint8(0), uint8(0)

	ans.Grow(max(len1, len2) + 1)
	for p1, p2 := len1-1, len2-1; p1 >= 0 || p2 >= 0; p1, p2 = p1-1, p2-1 {
		sum = 0
		if p1 >= 0 {
			sum = num1[p1] - '0'
		}
		if p2 >= 0 {
			sum += num2[p2] - '0'
		}
		sum += carry
		if sum > 9 {
			ans.WriteByte('0' + sum - 10)
			carry = 1
		} else {
			ans.WriteByte('0' + sum)
			carry = 0
		}
	}
	if carry == 1 {
		ans.WriteByte('1')
	}
	return reverse(ans.String())
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

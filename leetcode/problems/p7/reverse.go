package p7

import (
	"strconv"
)

const (
	max = (1 << 31) - 1
	min = -(1 << 31)
)

func reverse(x int) int {
	number := strconv.FormatInt(abs(x), 10)
	r := len(number) - 1
	for ; r >= 0 && number[r] == '0'; r-- {
	}
	result := int64(0)
	for base, l := int64(1), 0; l <= r; base, l = base*10, l+1 {
		result += int64(number[l]-'0') * base
	}

	if x < 0 {
		result *= -1
	}
	if result > max || result < min {
		return 0
	}
	return int(result)
}

func abs(n int) int64 {
	if n >= 0 {
		return int64(n)
	}
	return int64(-n)
}

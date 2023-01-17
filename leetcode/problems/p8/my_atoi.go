package p8

const max = 1<<31 - 1
const min = -(1 << 31)

func myAtoi(s string) int {
	pos, sign := 0, int64(1)
	for ; pos < len(s) && s[pos] == ' '; pos++ {
	}
	if pos == len(s) {
		return 0
	}
	if s[pos] == '-' || s[pos] == '+' {
		if s[pos] == '-' {
			sign = -1
		}
		pos++
	}
	for ; pos < len(s) && s[pos] == '0'; pos++ {
	}
	r := pos
	for ; r < len(s) && s[r] >= '0' && s[r] <= '9'; r++ {
	}
	if r-pos > 10 {
		if sign == 1 {
			return max
		} else {
			return min
		}
	}

	result := int64(0)
	for i, base := r-1, int64(1); i >= pos; i, base = i-1, base*10 {
		result += int64(s[i]-'0') * base
		if result*sign > max || result*sign < min {
			break
		}
	}

	result *= sign
	if result > max {
		return max
	}
	if result < min {
		return min
	}
	return int(result)
}

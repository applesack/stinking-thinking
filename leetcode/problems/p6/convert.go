package p6

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]strings.Builder, numRows)
	for flag, pos, i := 1, 0, 0; i < len(s); i++ {
		rows[pos].WriteByte(s[i])
		pos += flag
		if pos < 0 || pos >= numRows {
			flag = -flag
			pos += flag * 2
		}
	}
	builder := strings.Builder{}
	for i := 0; i < len(rows); i++ {
		builder.WriteString(rows[i].String())
	}
	return builder.String()
}

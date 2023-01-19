package p48

import (
	"fmt"
	"stinking-thinking/leetcode"
	"testing"
)

func Test48_1(t *testing.T) {
	mat := leetcode.Build2DSlice(
		leetcode.BuildSlice(1, 2, 3),
		leetcode.BuildSlice(4, 5, 6),
		leetcode.BuildSlice(7, 8, 9),
	)
	rotate(mat)
	fmt.Println(mat)
}

func Test48_2(t *testing.T) {
	mat := leetcode.Build2DSlice(
		leetcode.BuildSlice(1),
	)
	rotate(mat)
	fmt.Println(mat)
}

func Test48_3(t *testing.T) {
	mat := leetcode.Build2DSlice(
		leetcode.BuildSlice(5, 1, 9, 11),
		leetcode.BuildSlice(2, 4, 8, 10),
		leetcode.BuildSlice(13, 3, 6, 7),
		leetcode.BuildSlice(15, 14, 12, 16),
	)
	rotate(mat)
	fmt.Println(mat)
}

func Test48_4(t *testing.T) {
	mat := leetcode.Build2DSlice(
		leetcode.BuildSlice(1, 2),
		leetcode.BuildSlice(3, 4),
	)
	rotate(mat)
	fmt.Println(mat)
}

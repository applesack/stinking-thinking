package p240

import (
	"fmt"
	"stinking-thinking/leetcode"
	"testing"
)

func Test240_1(t *testing.T) {
	mat := leetcode.Build2DSlice(
		leetcode.BuildSlice(1, 4, 7, 11, 15),
		leetcode.BuildSlice(2, 5, 8, 12, 19),
		leetcode.BuildSlice(3, 6, 9, 16, 22),
		leetcode.BuildSlice(10, 13, 14, 17, 24),
		leetcode.BuildSlice(18, 21, 23, 26, 30),
	)
	fmt.Println(searchMatrix(mat, 20))
	fmt.Println(searchMatrix(mat, 10))
	fmt.Println(searchMatrix(mat, 13))
}

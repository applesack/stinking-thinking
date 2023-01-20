package p435

import (
	"fmt"
	"stinking-thinking/leetcode"
	"testing"
)

func Test435_1(t *testing.T) {
	mat := leetcode.Build2DSlice(
		leetcode.BuildSlice(1, 2),
		leetcode.BuildSlice(2, 3),
		leetcode.BuildSlice(3, 4),
		leetcode.BuildSlice(1, 3),
	)
	fmt.Println(eraseOverlapIntervals(mat))
}

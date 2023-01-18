package p56

import (
	"fmt"
	"stinking-thinking/leetcode"
	"testing"
)

func Test56_1(t *testing.T) {
	result := merge(leetcode.Build2DSlice(
		leetcode.BuildSlice(1, 3),
		leetcode.BuildSlice(2, 6),
		leetcode.BuildSlice(8, 10),
		leetcode.BuildSlice(15, 18),
	))
	fmt.Println(result)
}

func Test56_2(t *testing.T) {
	result := merge(leetcode.Build2DSlice(
		leetcode.BuildSlice(1, 4),
		leetcode.BuildSlice(4, 5),
	))
	fmt.Println(result)
}

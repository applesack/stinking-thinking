package p15

import (
	"fmt"
	"stinking-thinking/leetcode"
	"testing"
)

func Test15_1(t *testing.T) {
	ans := threeSum(leetcode.BuildSlice(-1, 0, 1, 2, -1, -4))
	fmt.Println(ans)
}

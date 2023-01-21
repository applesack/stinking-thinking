package p238

import (
	"fmt"
	"stinking-thinking/leetcode"
	"testing"
)

func Test238_1(t *testing.T) {
	fmt.Println(productExceptSelf(leetcode.BuildSlice(1, 2, 3, 4)))
	fmt.Println(productExceptSelf(leetcode.BuildSlice(-1, 1, 0, -3, 3)))
}

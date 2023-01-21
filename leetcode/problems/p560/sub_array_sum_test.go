package p560

import (
	"fmt"
	"stinking-thinking/leetcode"
	"testing"
)

func Test560_1(t *testing.T) {
	fmt.Println(subarraySum(leetcode.BuildSlice(1, 1, 1), 2))
	fmt.Println(subarraySum(leetcode.BuildSlice(1, 2, 3), 3))
	fmt.Println(subarraySum(leetcode.BuildSlice(), 3))
}

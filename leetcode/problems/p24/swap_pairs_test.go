package p24

import (
	"stinking-thinking/leetcode"
	"testing"
)

func Test24_1(t *testing.T) {
	list := leetcode.BuildListNode(1, 2, 3, 4)
	swapPairs(list).Display()
}

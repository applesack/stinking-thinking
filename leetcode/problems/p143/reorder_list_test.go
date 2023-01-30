package p143

import (
	"stinking-thinking/leetcode"
	"testing"
)

func Test143_1(t *testing.T) {
	head := leetcode.BuildListNode(1, 2, 3, 4)
	reorderList(head)
	head.Display()
}

func Test143_2(t *testing.T) {
	head := leetcode.BuildListNode(1, 2, 3, 4, 5)
	reorderList(head)
	head.Display()
}

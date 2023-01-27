package p82

import (
	"fmt"
	"stinking-thinking/leetcode"
	"testing"
)

func Test82_1(t *testing.T) {
	list := leetcode.BuildListNode(1, 2, 3, 3, 4, 4, 5)
	list = deleteDuplicates(list)
	fmt.Println()
}

func Test82_2(t *testing.T) {
	list := leetcode.BuildListNode(1, 1, 1, 2, 3, 3)
	list = deleteDuplicates(list)
	fmt.Println()
}

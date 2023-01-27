package p142

import (
	"fmt"
	"stinking-thinking/leetcode"
	"testing"
)

func Test142_1(t *testing.T) {
	n1, n2, n3, n4 :=
		leetcode.BuildListNode(3),
		leetcode.BuildListNode(2),
		leetcode.BuildListNode(0),
		leetcode.BuildListNode(-4)
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2
	fmt.Println(detectCycle(n2))
}

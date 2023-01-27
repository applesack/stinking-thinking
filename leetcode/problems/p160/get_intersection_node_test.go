package p160

import (
	"fmt"
	"stinking-thinking/leetcode"
	"testing"
)

func Test160_1(t *testing.T) {
	n1 := leetcode.BuildListNode(1)
	n9 := leetcode.BuildListNode(9)
	n11 := leetcode.BuildListNode(1)
	n2 := leetcode.BuildListNode(2)
	n4 := leetcode.BuildListNode(4)
	n3 := leetcode.BuildListNode(3)

	n1.Next = n9
	n9.Next = n11
	n11.Next = n2
	n2.Next = n4
	n3.Next = n2

	rsl := getIntersectionNode(n1, n3)
	fmt.Println(rsl)
}

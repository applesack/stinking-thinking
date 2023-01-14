package p2

import (
	"fmt"
	. "stinking-thinking/leetcode"
	"testing"
)

func Test2_1(t *testing.T) {
	l1 := BuildListNode(2, 4, 3)
	l2 := BuildListNode(5, 6, 4)
	rsl := addTwoNumbers(l1, l2)
	rsl.Display()
}

func Test2_2(t *testing.T) {
	l1 := BuildListNode(9, 9, 9, 9, 9, 9, 9)
	l2 := BuildListNode(9, 9, 9, 9)
	rsl := addTwoNumbers(l1, l2)
	fmt.Println(rsl)
}

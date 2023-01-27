package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (t *ListNode) String() string {
	return fmt.Sprintf("{Val:%d,Next:%d}", t.Val, t.Next)
}

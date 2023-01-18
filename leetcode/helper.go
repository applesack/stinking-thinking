package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func BuildListNode(nums ...int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	var point, head *ListNode
	head = &ListNode{}
	point = head

	for _, num := range nums {
		node := &ListNode{Val: num}
		point.Next = node
		point = point.Next
	}

	return head.Next
}

func (t *ListNode) Display() {
	if t == nil {
		fmt.Println("nil")
	}
	p := t
	buff := strings.Builder{}
	for p != nil {
		value := p.Val
		buff.WriteString(strconv.Itoa(value))
		p = p.Next
		if p != nil {
			buff.WriteString(" -> ")
		}
	}
	fmt.Println(buff.String())
}

func BuildSlice(nums ...int) []int {
	slice := make([]int, len(nums))
	for i, num := range nums {
		slice[i] = num
	}
	return slice
}

func Build2DSlice(slices ...[]int) [][]int {
	slice2d := make([][]int, len(slices))
	for i, slice := range slices {
		slice2d[i] = slice
	}
	return slice2d
}

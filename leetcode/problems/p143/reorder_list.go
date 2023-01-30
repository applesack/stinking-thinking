package p143

import . "stinking-thinking/leetcode"

func reorderList(head *ListNode) {
	slice := store(head)
	vHead, p := new(ListNode), head
	p = vHead
	for l, r := 0, len(slice)-1; l <= r; l, r = l+1, r-1 {
		p.Next = slice[l]
		p = p.Next
		if l != r {
			p.Next = slice[r]
			p = p.Next
		}
		p.Next = nil
	}
}

func store(head *ListNode) []*ListNode {
	slice := make([]*ListNode, 0, 4)
	for p := head; p != nil; p = p.Next {
		slice = append(slice, p)
	}
	return slice
}

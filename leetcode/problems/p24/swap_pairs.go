package p24

import . "stinking-thinking/leetcode"

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	vHead := new(ListNode)
	pre, cur, nxt, tail := vHead, head, head.Next, vHead
	vHead.Next = head
	for cur != nil && nxt != nil {
		tail = nxt.Next
		pre.Next = nxt
		nxt.Next = cur
		cur.Next = tail

		cur, nxt = nxt, cur
		pre, cur, nxt = pre.Next, cur.Next, nxt.Next
		if nxt != nil {
			pre, cur, nxt = pre.Next, cur.Next, nxt.Next
		}
	}

	return vHead.Next
}

package p82

import . "stinking-thinking/leetcode"

func deleteDuplicates(head *ListNode) *ListNode {
	vHead := new(ListNode)
	vHead.Next = head
	pre, cur := vHead, head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			value := cur.Val
			for cur != nil && cur.Val == value {
				cur = cur.Next
			}
			pre.Next = cur
			continue
		}
		pre = cur
		cur = cur.Next
	}
	return vHead.Next
}

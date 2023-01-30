package p25

import . "stinking-thinking/leetcode"

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	count := getNodeCount(head)
	vHead := new(ListNode)
	pre, cur, tmp := vHead, head, head
	vp, vh, limit, pos := head, head, 0, 0
	pre.Next = cur
	for {
		if cur == nil || pos+k > count {
			break
		}
		vp, vh = pre, cur
		for limit = pos + k; pos < limit; pos++ {
			tmp = pre
			pre, cur = cur, cur.Next
			pre.Next = tmp
		}
		vp.Next = pre
		vh.Next = cur
		pre = vh
	}
	return vHead.Next
}

func getNodeCount(head *ListNode) int {
	p, count := head, 0
	for p != nil {
		p = p.Next
		count++
	}
	return count
}

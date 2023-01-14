package p2

import . "stinking-thinking/leetcode"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		if l1 == nil {
			return l2
		}
		return l1
	}

	var head, point *ListNode
	var carry = 0
	head = &ListNode{}
	point = head

	// 两个链表向前推进, 每次对并行的两个节点的值进行相加
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		next := &ListNode{Val: sum % 10}
		point.Next = next
		point = point.Next
		if sum > 9 {
			carry = 1
		} else {
			carry = 0
		}
	}

	if carry != 0 {
		point.Next = &ListNode{Val: carry}
	}

	return head.Next
}

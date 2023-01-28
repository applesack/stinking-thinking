package p707

type node struct {
	val        int
	prev, next *node
}

type MyLinkedList struct {
	head, tail *node
	size       int
}

func Constructor() MyLinkedList {
	h, t := new(node), new(node)
	bind(h, t)
	return MyLinkedList{
		head: h,
		tail: t,
	}
}

func (t *MyLinkedList) Get(index int) int {
	if index < 0 || index >= t.size {
		return -1
	}
	p := t.head
	for i := 0; i <= index; i++ {
		p = p.next
	}
	return p.val
}

func (t *MyLinkedList) AddAtHead(val int) {
	n := newNode(val)
	bind(n, t.head.next)
	bind(t.head, n)
	t.size++
}

func (t *MyLinkedList) AddAtTail(val int) {
	n := newNode(val)
	bind(t.tail.prev, n)
	bind(n, t.tail)
	t.size++
}

func (t *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		index = 0
	}
	if index > t.size {
		return
	}
	pre, cur := t.head, t.head.next
	for i := 0; i < t.size && i < index; i++ {
		pre, cur = pre.next, cur.next
	}
	n := newNode(val)
	bind(n, cur)
	bind(pre, n)
	t.size++
}

func (t *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= t.size {
		return
	}
	p := t.head.next
	for i := 0; i < index; i++ {
		p = p.next
	}
	bind(p.prev, p.next)
	t.size--
}

func newNode(value int) *node {
	n := new(node)
	n.val = value
	return n
}

func bind(p, n *node) {
	p.next = n
	n.prev = p
}

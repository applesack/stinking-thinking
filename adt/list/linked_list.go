package list

import (
	"fmt"
)

type LinkedList[T any] struct {
	size int
	head *doublyNode[T]
	tail *doublyNode[T]
}

type doublyNode[T any] struct {
	value T
	prev  *doublyNode[T]
	next  *doublyNode[T]
}

func CreateLinkedList[T any]() *LinkedList[T] {
	h, t := &doublyNode[T]{}, &doublyNode[T]{}
	bindDoublyNode(h, t)
	return &LinkedList[T]{size: 0, head: h, tail: t}
}

func (t *LinkedList[T]) Size() int {
	return t.size
}

func (t *LinkedList[T]) AddLast(value T) {
	prev, tail, node := t.tail.prev, t.tail, createDoublyNode(value)
	bindDoublyNode(prev, node)
	bindDoublyNode(node, tail)
	t.size++
}

func (t *LinkedList[T]) AddFirst(value T) {
	head, next, node := t.head, t.head.next, createDoublyNode(value)
	bindDoublyNode(head, node)
	bindDoublyNode(node, next)
	t.size++
}

func (t *LinkedList[T]) Get(index int) (T, error) {
	if err := t.checkIndex(index); err != nil {
		return t.head.value, err
	}
	return t.locateIndex(index).value, nil
}

func (t *LinkedList[T]) Remove(index int) (T, error) {
	if err := t.checkIndex(index); err != nil {
		return t.head.value, err
	}
	ptr := t.locateIndex(index)
	p, n := ptr.prev, ptr.next
	bindDoublyNode(p, n)
	t.size--
	return ptr.value, nil
}

func (t *LinkedList[T]) RemoveFirst() (T, error) {
	return t.Remove(0)
}

func (t *LinkedList[T]) RemoveLast() (T, error) {
	return t.Remove(t.size - 1)
}

func (t *LinkedList[T]) Values() []T {
	arr, pos := make([]T, t.size), 0
	for ptr := t.head.next; ptr != t.tail; ptr, pos = ptr.next, pos+1 {
		arr[pos] = ptr.value
	}
	return arr
}

func (t *LinkedList[T]) GetFirst() (bool, T) {
	if t.size == 0 {
		return false, t.head.value
	}
	return true, t.head.next.value
}

func (t *LinkedList[T]) GetLast() (bool, T) {
	if t.size == 0 {
		return false, t.tail.value
	}
	return true, t.tail.prev.value
}

func (t *LinkedList[T]) locateIndex(index int) *doublyNode[T] {
	ptr, pos := t.head, 0
	if t.size-index > index {
		for ptr, pos = t.head, 0; pos <= index; ptr, pos = ptr.next, pos+1 {
		}
	} else {
		for ptr, pos = t.tail, t.size; pos > index; ptr, pos = ptr.prev, pos-1 {
		}
	}
	return ptr
}

func (t *LinkedList[T]) checkIndex(index int) error {
	if index < 0 || index >= t.size {
		return fmt.Errorf("index out of bound [%d:%d]", t.size, index)
	}
	return nil
}

func bindDoublyNode[T any](prev, next *doublyNode[T]) {
	prev.next = next
	next.prev = prev
}

func createDoublyNode[T any](value T) *doublyNode[T] {
	return &doublyNode[T]{value: value}
}

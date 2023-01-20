package list

import (
	"container/list"
	"fmt"
	. "stinking-thinking/adt"
)

type LinkedList[T any] struct {
	size int
	head *DLinkedNode[T]
	tail *DLinkedNode[T]
}

func CreateLinkedList[T any]() *LinkedList[T] {
	h, t := &DLinkedNode[T]{}, &DLinkedNode[T]{}
	BindDLinkedNode(h, t)
	return &LinkedList[T]{size: 0, head: h, tail: t}
}

func (t *LinkedList[T]) Size() int {
	return t.size
}

func (t *LinkedList[T]) RPush(value T) {
	prev, tail, node := t.tail.Prev, t.tail, CreateDLinkedNode(value)
	BindDLinkedNode(prev, node)
	BindDLinkedNode(node, tail)
	t.size++
}

func (t *LinkedList[T]) LPush(value T) {
	head, next, node := t.head, t.head.Next, CreateDLinkedNode(value)
	BindDLinkedNode(head, node)
	BindDLinkedNode(node, next)
	t.size++
}

func (t *LinkedList[T]) Get(index int) (*T, error) {
	list.New()
	if err := t.checkIndex(index); err != nil {
		return nil, err
	}
	return &(t.locateIndex(index).Value), nil
}

func (t *LinkedList[T]) Remove(index int) (*T, error) {
	if err := t.checkIndex(index); err != nil {
		return nil, err
	}
	ptr := t.locateIndex(index)
	ptr.Delete()
	t.size--
	return &ptr.Value, nil
}

func (t *LinkedList[T]) LPop() (*T, error) {
	return t.Remove(0)
}

func (t *LinkedList[T]) RPop() (*T, error) {
	return t.Remove(t.size - 1)
}

func (t *LinkedList[T]) Values() []T {
	arr, pos := make([]T, t.size), 0
	for ptr := t.head.Next; ptr != t.tail; ptr, pos = ptr.Next, pos+1 {
		arr[pos] = ptr.Value
	}
	return arr
}

func (t *LinkedList[T]) LGet() (*T, bool) {
	if t.size == 0 {
		return nil, false
	}
	return &t.head.Next.Value, true
}

func (t *LinkedList[T]) RGet() (*T, bool) {
	if t.size == 0 {
		return nil, false
	}
	return &(t.tail.Prev.Value), true
}

func (t *LinkedList[T]) locateIndex(index int) *DLinkedNode[T] {
	ptr, pos := t.head, 0
	if t.size-index > index {
		for ptr, pos = t.head, 0; pos <= index; ptr, pos = ptr.Next, pos+1 {
		}
	} else {
		for ptr, pos = t.tail, t.size; pos > index; ptr, pos = ptr.Prev, pos-1 {
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

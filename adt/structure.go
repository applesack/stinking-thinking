package adt

import "fmt"

type DLinkedNode[T any] struct {
	Value T
	Prev  *DLinkedNode[T]
	Next  *DLinkedNode[T]
}

type Entry[K any, V any] struct {
	Key   K
	Value V
}

func NewDLinkedNode[T any](value T) *DLinkedNode[T] {
	return &DLinkedNode[T]{Value: value}
}

func NewEntry[K any, V any](key K, value V) *Entry[K, V] {
	return &Entry[K, V]{Key: key, Value: value}
}

func BindDLinkedNode[T any](prev, next *DLinkedNode[T]) {
	prev.Next = next
	next.Prev = prev
}

func (t *DLinkedNode[T]) Delete() {
	p, n := t.Prev, t.Next
	BindDLinkedNode(p, n)
	t.Prev, t.Next = nil, nil
}

func (t *DLinkedNode[T]) String() string {
	return fmt.Sprintf("[v=%+v, p=%b, n=%b]", t.Value, t.Prev, t.Next)
}

func (t *Entry[K, V]) String() string {
	return fmt.Sprintf("[k=%v, v=%v]", t.Key, t.Value)
}

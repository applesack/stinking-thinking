package tree

import (
	. "stinking-thinking/adt"
)

type LRU[K comparable, V any] struct {
	maxSize    int
	m          map[K]*DLinkedNode[*Entry[K, V]]
	head, tail *DLinkedNode[*Entry[K, V]]
}

func CreateLRU[K comparable, V any](maxSize int) *LRU[K, V] {
	h, t := &DLinkedNode[*Entry[K, V]]{}, &DLinkedNode[*Entry[K, V]]{}
	BindDLinkedNode(h, t)
	return &LRU[K, V]{
		maxSize: maxSize,
		m:       make(map[K]*DLinkedNode[*Entry[K, V]]),
		head:    h,
		tail:    t,
	}
}

func (t *LRU[K, V]) Size() int {
	return len(t.m)
}

func (t *LRU[K, V]) Get(key K) (*V, bool) {
	previous, exists := t.m[key]
	if exists {
		t.moveToHead(previous)
		return &previous.Value.Value, true
	}
	return nil, false
}

func (t *LRU[K, V]) Put(key K, value V) (*V, bool) {
	previous, exists := t.m[key]
	if exists {
		previousValue := previous.Value.Value
		previous.Value.Value = value
		t.moveToHead(previous)
		return &previousValue, true
	}

	entry := CreateEntry(key, value)
	node := CreateDLinkedNode(entry)
	if t.Size() >= t.maxSize {
		t.removeTail()
	}
	t.m[key] = node
	t.addToHead(node)
	return nil, false
}

func (t *LRU[K, V]) Remove(key K) (*V, bool) {
	previous, exists := t.m[key]
	if exists {
		delete(t.m, key)
		previous.Delete()
		return &previous.Value.Value, true
	}
	return nil, false
}

func (t *LRU[K, V]) Entries() []*Entry[K, V] {
	slice := make([]*Entry[K, V], t.Size())
	for ptr, pos := t.head.Next, 0; ptr != t.tail; ptr, pos = ptr.Next, pos+1 {
		slice[pos] = ptr.Value
	}
	return slice
}

func (t *LRU[K, V]) moveToHead(node *DLinkedNode[*Entry[K, V]]) {
	node.Delete()
	t.addToHead(node)
}

func (t *LRU[K, V]) addToHead(node *DLinkedNode[*Entry[K, V]]) {
	BindDLinkedNode(node, t.head.Next)
	BindDLinkedNode(t.head, node)
}

func (t *LRU[K, V]) removeTail() {
	eldest := t.tail.Prev
	eldest.Delete()
	delete(t.m, eldest.Value.Key)
}

package p706

import (
	"fmt"
	"testing"
)

func Test706_1(t *testing.T) {
	m := Constructor()
	m.Put(1, 1)
	m.Put(2, 2)
	fmt.Println(m.Get(1))
	fmt.Println(m.Get(3))
	m.Put(2, 1)
	fmt.Println(m.Get(2))
	m.Remove(2)
	fmt.Println(m.Get(2))
}

func Test706_2(t *testing.T) {
	m := Constructor()
	m.Put(1, 1)
	m.Put(2, 1)
	m.Put(3, 1)
	m.Put(4, 1)
	m.Put(5, 1)
	m.Put(6, 1)
	m.Put(7, 1)
	m.Put(8, 1)
}

func Test706_3(t *testing.T) {
	m := Constructor()
	m.Put(0, 0)
	m.Put(8, 8)
	fmt.Println(m.Get(0))
	fmt.Println(m.Get(8))
}

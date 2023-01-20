package list

import (
	"fmt"
	"testing"
)

func TestLinkedList_AddFirst(t *testing.T) {
	list := CreateLinkedList[int]()
	list.AddLast(1)
	list.AddLast(2)
	list.AddFirst(3)
	fmt.Println(list.Size())
	fmt.Println(list.Values())
}

func TestLinkedList_Get(t *testing.T) {
	list := CreateLinkedList[int]()
	for i := 0; i < 10; i++ {
		list.AddLast(i)
	}
	for i := -2; i < 12; i++ {
		value, err := list.Get(i)
		if err != nil {
			fmt.Println(fmt.Sprintf("i=%d, value=%s", i, err.Error()))
		} else {
			fmt.Println(fmt.Sprintf("i=%d, value=%d", i, value))
		}
	}
}

func TestLinkedList_Remove(t *testing.T) {
	displayRm := func(ls *LinkedList[int], idx int) {
		value, err := ls.Remove(idx)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(fmt.Sprintf("removed %d", value))
			fmt.Println(ls.Values())
		}
	}

	list := CreateLinkedList[int]()
	for i := 0; i < 10; i++ {
		list.AddLast(i)
	}
	displayRm(list, 3)
}

func TestLinkedList_RemoveBoth(t *testing.T) {
	list := CreateLinkedList[int]()
	for i := 0; i < 10; i++ {
		list.AddLast(i)
	}
	_, _ = list.RemoveFirst()
	_, _ = list.RemoveLast()
	fmt.Println(list.Values())
}

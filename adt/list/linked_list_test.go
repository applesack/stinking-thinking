package list

import (
	"fmt"
	"strconv"
	"testing"
)

func TestLinkedList_AddFirst(t *testing.T) {
	list := CreateLinkedList[int]()
	list.RPush(1)
	list.RPush(2)
	list.LPush(3)
	fmt.Println(list.Size())
	fmt.Println(list.Values())
}

func TestLinkedList_Get(t *testing.T) {
	list := CreateLinkedList[int]()
	for i := 0; i < 10; i++ {
		list.RPush(i)
	}
	for i := -2; i < 12; i++ {
		value, err := list.Get(i)
		if err != nil {
			fmt.Println(fmt.Sprintf("i=%d, value=%s", i, err.Error()))
		} else {
			fmt.Println(fmt.Sprintf("i=%d, value=%d", i, *value))
		}
	}
}

func TestLinkedList_Remove(t *testing.T) {
	displayRm := func(ls *LinkedList[int], idx int) {
		value, err := ls.Remove(idx)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(fmt.Sprintf("removed %d", *value))
			fmt.Println(ls.Values())
		}
	}

	list := CreateLinkedList[int]()
	for i := 0; i < 10; i++ {
		list.RPush(i)
	}
	displayRm(list, 3)
}

func TestLinkedList_RemoveBoth(t *testing.T) {
	list := CreateLinkedList[int]()
	for i := 0; i < 10; i++ {
		list.RPush(i)
	}
	_, _ = list.LPop()
	_, _ = list.RPop()
	fmt.Println(list.Values())
}

func TestLinkedList_Struct(t *testing.T) {
	list := CreateLinkedList[testStruct]()
	has, s := list.LGet()
	fmt.Println(has, s)
}

func BenchmarkLinkedList_Batch(b *testing.B) {
	b.StartTimer()
	defer b.StopTimer()
	l := CreateLinkedList[int]()
	for i := 0; i < 100000; i++ {
		l.RPush(i)
	}
	for i := 0; i < 100000; i++ {
		_, _ = l.Get(i)
	}
}

func BenchmarkLinkedList_Batch1(b *testing.B) {
	b.StartTimer()
	defer b.StopTimer()
	l := CreateLinkedList[testStruct]()
	for i := 0; i < 100000; i++ {
		l.RPush(testStruct{
			name: strconv.Itoa(i),
			age:  i,
		})
	}
	for i := 0; i < 100000; i++ {
		_, _ = l.Get(i)
	}
}

type testStruct struct {
	name string
	age  int
}

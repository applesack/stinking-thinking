package p707

import (
	"fmt"
	"testing"
)

func Test707_1(t *testing.T) {
	list := Constructor()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)
	fmt.Println(list.Get(1))
	list.DeleteAtIndex(1)
	fmt.Println(list.Get(1))
}

func Test707_2(t *testing.T) {
	list := Constructor()
	list.AddAtHead(1)
	list.DeleteAtIndex(0)
	fmt.Println(list)
}

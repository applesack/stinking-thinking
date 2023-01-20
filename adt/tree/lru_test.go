package tree

import (
	"fmt"
	"testing"
)

func TestLRU_PUT_GET(t *testing.T) {
	cache := CreateLRU[int, int](3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	fmt.Println(cache.Size())
	fmt.Println(cache.Entries())
}

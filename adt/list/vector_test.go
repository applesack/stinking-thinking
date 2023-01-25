package list

import (
	"fmt"
	"testing"
)

func TestVector_Add(t *testing.T) {
	v := NewVector[int](1)
	for i := 0; i < 10; i++ {
		v.Add(i)
	}
	fmt.Println(v)
}

// 10000 - 0.09264
// 1000  - 0.1329
func BenchmarkVector_Add(b *testing.B) {
	b.StartTimer()
	defer b.StopTimer()

	for i := 0; i < 1000; i++ {
		ls := NewVector[int](2048)
		for i := 0; i < 10000; i++ {
			ls.Add(i)
		}
		for i := 0; i < 10000; i++ {
			_, _ = ls.Get(i)
		}
	}
}

// 10000 - 0.06375
// 1000  - 0.004996
func BenchmarkSlice(b *testing.B) {
	b.StartTimer()
	defer b.StopTimer()
	for i := 0; i < 1000; i++ {
		ls := make([]int, 0)
		for i := 0; i < 10000; i++ {
			ls = append(ls, i)
		}
		for i := 0; i < 10000; i++ {
			_ = ls[i]
		}
	}
}

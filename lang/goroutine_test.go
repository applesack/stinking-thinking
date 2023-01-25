package lang

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func TestGoroutine(t *testing.T) {
	var wg = sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go goroutine(&wg)
	}
	wg.Wait()
	fmt.Println("ok " + strconv.Itoa(i))
}

var i = 0
var lock sync.Mutex

func goroutine(wg *sync.WaitGroup) {
	defer wg.Done()
	lock.Lock()
	i++
	lock.Unlock()
}

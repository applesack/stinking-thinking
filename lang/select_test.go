package lang

import (
	"fmt"
	"sync"
	"testing"
)

var ch = make(chan int, 1024)
var chClose = make(chan bool)

func TestSelect_1(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		tmp := i
		go func() {
			defer wg.Done()
			asyncTask(tmp)
		}()
	}
	asyncReceive()
	wg.Wait()

}

func asyncTask(i int) {
	ch <- i
}

func asyncReceive() {
LOOP:
	for {
		select {
		case num := <-ch:
			fmt.Println(num)
		case <-chClose:
			close(ch)
			close(chClose)
			break LOOP
		default:
		}
	}
	fmt.Println("通道已关闭")
}

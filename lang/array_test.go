package lang

import (
	"fmt"
	"testing"
)

func TestRef(t *testing.T) {
	// 公共方法
	var arr = []int{1, 2}
	modify(&arr)
	fmt.Println(arr)
}

func modify(arr *[]int) {
	(*arr)[0] = 4
}

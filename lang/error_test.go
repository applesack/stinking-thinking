package lang

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_1(t *testing.T) {
	num, err := errorFun(0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
}

func errorFun(n int) (int, error) {
	if n == 0 {
		return n, nil
	}
	return 0, errors.New("test")
}

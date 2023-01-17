package p7

import (
	"fmt"
	"testing"
)

func Test7_1(t *testing.T) {
	fmt.Println(reverse(123))
}

func Test7_2(t *testing.T) {
	fmt.Println(reverse(-123))
}

func Test7_3(t *testing.T) {
	fmt.Println(reverse(0))
}

func Test7_4(t *testing.T) {
	fmt.Println(reverse(1534236469))
}

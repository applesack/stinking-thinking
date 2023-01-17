package p8

import (
	"fmt"
	"testing"
)

func Test8_1(t *testing.T) {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("    -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("+1"))
	fmt.Println(myAtoi("9223372036854775808"))
}

func Test8_2(t *testing.T) {
	text := "100000000000000000000000" +
		"000000000000000000000000000000" +
		"000000000000000000000000000000" +
		"00000000000000000000522545459"
	fmt.Println(myAtoi(text))
}

package lang

import (
	"fmt"
	"testing"
)

func TestString_1(t *testing.T) {
	str := "0123456789"
	str1 := str[1:5]
	str2 := str[4:]
	fmt.Println(str1)
	fmt.Println(str2)
}

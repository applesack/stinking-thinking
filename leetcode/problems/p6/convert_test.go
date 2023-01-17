package p6

import (
	"fmt"
	"testing"
)

func Test6_1(t *testing.T) {
	rsl := convert("PAYPALISHIRING", 3)
	fmt.Println(rsl)
}

func Test6_2(t *testing.T) {
	rsl := convert("PAYPALISHIRING", 1)
	fmt.Println(rsl)
}

func Test6_3(t *testing.T) {
	rsl := convert("PAYPALISHIRING", 4)
	fmt.Println(rsl)
}

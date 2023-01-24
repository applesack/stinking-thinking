package p43

import (
	"fmt"
	"stinking-thinking/leetcode"
	"testing"
)

func Test43_1(t *testing.T) {
	fmt.Println(multiply("2", "5"))
	fmt.Println(multiply("123", "456"))
	fmt.Println(multiply("99", "99"))
	fmt.Println(multiply("99", "10"))
	fmt.Println(multiply("9133", "0"))
}

func Test43_2(t *testing.T) {
	fmt.Println(multiply("140", "721"))
}

func TestMultiple2(t *testing.T) {
	fmt.Println(multiply2(leetcode.BuildSlice[uint8](9, 9), 9, 0))
	fmt.Println(multiply2(leetcode.BuildSlice[uint8](9, 9, 9, 9), 9, 0))
	fmt.Println(multiply2(leetcode.BuildSlice[uint8](9, 9, 9, 9), 9, 1))
}

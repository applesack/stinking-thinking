package p3

import (
	"fmt"
	"testing"
)

func Test3_0(t *testing.T) {
	fmt.Println(int('a'))
	fmt.Println(int('_'))
	fmt.Println(int('+'))
	fmt.Println(int('A'))
}

func Test3_1(t *testing.T) {
	length := lengthOfLongestSubstring("")
	fmt.Println(length)
}

func Test3_2(t *testing.T) {
	length := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(length)
}

func Test3_3(t *testing.T) {
	length := lengthOfLongestSubstring("bbbbb")
	fmt.Println(length)
}

func Test3_4(t *testing.T) {
	length := lengthOfLongestSubstring("pwwkew")
	fmt.Println(length)
}

package p5

import (
	"fmt"
	"testing"
)

func Test5_1(t *testing.T) {
	fmt.Println(longestPalindrome("aba"))
}

func Test5_2(t *testing.T) {
	fmt.Println(longestPalindrome("abaa"))
}

func Test5_3(t *testing.T) {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}

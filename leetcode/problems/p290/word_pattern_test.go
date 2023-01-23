package p290

import (
	"fmt"
	"testing"
)

func Test290_1(t *testing.T) {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat dogs"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))
}

package p49

import (
	"fmt"
	"stinking-thinking/leetcode"
	"testing"
)

func Test49_1(t *testing.T) {
	fmt.Println(groupAnagrams(leetcode.BuildSlice("eat", "tea", "tan", "ate", "nat", "bat")))
	fmt.Println(groupAnagrams(leetcode.BuildSlice("")))
	fmt.Println(groupAnagrams(leetcode.BuildSlice("a")))
}

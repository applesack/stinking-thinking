package p1249

import (
	"strings"
)

func minRemoveToMakeValid(s string) string {
	left := make([]int, 0, 4)
	indexToRemove := make(map[int]bool)
	for i, ch := range s {
		if ch == ')' {
			if len(left) == 0 {
				indexToRemove[i] = true
			} else {
				left = pop(left)
			}
		} else if ch == '(' {
			left = push(left, i)
		}
	}
	for _, i := range left {
		indexToRemove[i] = true
	}
	buff := new(strings.Builder)
	for i, ch := range s {
		if _, ok := indexToRemove[i]; ok {
			continue
		} else {
			buff.WriteString(string(ch))
		}
	}

	return buff.String()
}

func push(nums []int, nun int) []int {
	return append(nums, nun)
}

func pop(nums []int) []int {
	return nums[0 : len(nums)-1]
}

package p25

import (
	"stinking-thinking/leetcode"
	"testing"
)

func Test25_1(t *testing.T) {
	reverseKGroup(leetcode.BuildListNode(1, 2, 3, 4, 5), 2).Display()
}

func Test25_2(t *testing.T) {
	reverseKGroup(leetcode.BuildListNode(1, 2, 3, 4, 5), 3).Display()
}

func Test25_3(t *testing.T) {
	reverseKGroup(leetcode.BuildListNode(1, 2, 3, 4, 5, 6, 7, 8, 9), 3).Display()
}

func Test25_4(t *testing.T) {
	reverseKGroup(leetcode.BuildListNode(1, 2, 3, 4, 5, 6), 2).Display()
}

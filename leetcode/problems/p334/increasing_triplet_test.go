package p334

import (
	"fmt"
	"stinking-thinking/leetcode"
	"testing"
)

func Test334_1(t *testing.T) {
	fmt.Println(increasingTriplet(leetcode.BuildSlice(1, 2, 3, 4, 5)))
	fmt.Println(increasingTriplet(leetcode.BuildSlice(5, 4, 3, 2, 1)))
	fmt.Println(increasingTriplet(leetcode.BuildSlice(2, 1, 5, 0, 4, 6)))
}

func Test334_2(t *testing.T) {
	fmt.Println(increasingTriplet(leetcode.BuildSlice(20, 100, 10, 12, 5, 13)))
}

func Test334_3(t *testing.T) {
	fmt.Println(increasingTriplet(leetcode.BuildSlice(1, 1, 1, 1, 1, 1, 1, 1, 3, 7)))
}

func Test334_4(t *testing.T) {
	fmt.Println(increasingTriplet(leetcode.BuildSlice(1, 5, 0, 4, 1, 3)))
}

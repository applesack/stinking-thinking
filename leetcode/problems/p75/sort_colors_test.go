package p75

import (
	"fmt"
	"testing"
)

func Test75_1(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func Test75_2(t *testing.T) {
	nums := []int{2, 2, 1, 1}
	sortColors(nums)
	fmt.Println(nums)
}

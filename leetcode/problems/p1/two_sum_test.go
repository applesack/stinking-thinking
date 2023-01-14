package p1

import (
	"fmt"
	"testing"
)

func Test1_1(t *testing.T) {
	arr := []int{2, 7, 11, 15}
	rsl := twoSum(arr, 9)
	fmt.Println(rsl)
}

func Test1_2(t *testing.T) {
	arr := []int{3, 2, 4}
	rsl := twoSum(arr, 6)
	fmt.Println(rsl)
}

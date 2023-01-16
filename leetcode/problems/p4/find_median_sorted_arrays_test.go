package p4

import (
	"fmt"
	"testing"
)

func Test4_1(t *testing.T) {
	arr1 := []int{1, 3}
	arr2 := []int{2}
	rsl := findMedianSortedArrays(arr1, arr2)
	fmt.Println(rsl)
}

func Test4_2(t *testing.T) {
	arr1 := []int{1, 2}
	arr2 := []int{3, 4}
	rsl := findMedianSortedArrays(arr1, arr2)
	fmt.Println(rsl)
}

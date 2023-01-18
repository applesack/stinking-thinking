package lang

import (
	"fmt"
	"sort"
	"testing"
)

func TestSlice_1(t *testing.T) {
	var nums []int
	fmt.Println(nums)
}

func TestSlice_2(t *testing.T) {
	var nums = []int{4, 2, 3, 5, 1}
	sortArr(nums)
	fmt.Println(nums)
}

func TestSlice_3(t *testing.T) {
	var nums = []int{4, 2, 3, 5, 1}
	modifyArr(nums)
	fmt.Println(nums)
}

func TestSlice_4(t *testing.T) {
	var nums = []int{1, 4, 1, 2, 5, 3, 3, 5, 1, 2}
	nums1 := copyArr(nums)
	fmt.Println(nums1)
}

func sortArr(nums []int) {
	sort.Ints(nums)
}

func modifyArr(nums []int) {
	nums[2] = 9
}

func copyArr(nums []int) []int {
	sort.Ints(nums)
	var result []int
	for i := 0; i < len(nums); i++ {
		result = append(result, nums[i])
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
	}
	return result
}

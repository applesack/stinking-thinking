package p15

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var ans [][]int
	if len(nums) < 3 {
		return ans
	}

	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r, sum := i+1, len(nums)-1, 0
		for l < r {
			sum = nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l, r = l+1, r-1
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return ans
}

package p334

import "math"

func increasingTriplet(nums []int) bool {
	first, second := nums[0], math.MaxInt
	for i := 1; i < len(nums); i++ {
		third := nums[i]
		if third > second {
			return true
		} else if third > first {
			second = third
		} else {
			first = third
		}
	}
	return false
}

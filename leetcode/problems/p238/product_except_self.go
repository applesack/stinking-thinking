package p238

func productExceptSelf(nums []int) []int {
	L, R := make([]int, len(nums)), make([]int, len(nums))

	R[len(nums)-1] = 1
	for i := len(R) - 2; i >= 0; i-- {
		R[i] = nums[i+1] * R[i+1]
	}

	L[0] = 1
	for i := 1; i < len(nums); i++ {
		L[i] = nums[i-1] * L[i-1]
	}

	for i := range nums {
		nums[i] = L[i] * R[i]
	}

	return nums
}

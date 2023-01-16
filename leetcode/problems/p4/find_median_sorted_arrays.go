package p4

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := make([]int, len(nums1)+len(nums2))
	limit1, limit2 := len(nums1), len(nums2)
	pos, pos1, pos2 := 0, 0, 0
	for pos1 < limit1 && pos2 < limit2 {
		if nums1[pos1] < nums2[pos2] {
			nums[pos] = nums1[pos1]
			pos1++
		} else {
			nums[pos] = nums2[pos2]
			pos2++
		}
		pos++
	}
	for ; pos1 < limit1; pos++ {
		nums[pos] = nums1[pos1]
		pos1++
	}
	for ; pos2 < limit2; pos++ {
		nums[pos] = nums2[pos2]
		pos2++
	}
	sum := len(nums)
	if (sum/2)*2 == sum {
		return (float64(nums[sum/2]) + float64(nums[sum/2-1])) / 2
	}
	return float64(nums[sum/2])
}

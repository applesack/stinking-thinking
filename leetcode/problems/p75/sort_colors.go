package p75

func sortColors(nums []int) {
	r, w, b, pos := 0, 0, 0, 0
	for _, num := range nums {
		switch num {
		case 0:
			r++
		case 1:
			w++
		case 2:
			b++
		}
	}
	for ; pos < r; pos++ {
		nums[pos] = 0
	}
	for ; pos < (r + w); pos++ {
		nums[pos] = 1
	}
	for ; pos < len(nums); pos++ {
		nums[pos] = 2
	}
}

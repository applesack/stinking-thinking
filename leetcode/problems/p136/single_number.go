package p136

// A^A=0
// A^B^A^A=0
// 0^A=A
func singleNumber(nums []int) int {
	var ans = nums[0]
	for i := 1; i < len(nums); i++ {
		ans = ans ^ nums[i]
	}
	return ans
}

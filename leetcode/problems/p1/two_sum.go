package p1

func twoSum(nums []int, target int) []int {
	for x := 0; x < len(nums); x++ {
		for y := 0; y < len(nums); y++ {
			if x != y && nums[x]+nums[y] == target {
				return []int{x, y}
			}
		}
	}
	return []int{}
}

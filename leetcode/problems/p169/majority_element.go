package p169

func majorityElement(nums []int) int {
	ans, cur, cnt := nums[0], 0, 1
	for i := 1; i < len(nums); i++ {
		cur = nums[i]
		if cur == ans {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			ans = cur
			cnt = 1
		}
	}
	return ans
}

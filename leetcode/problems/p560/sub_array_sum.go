package p560

func subarraySum(nums []int, k int) int {
	pre, m, ans := 0, make(map[int]int), 0
	m[0] = 1
	for _, num := range nums {
		pre += num
		tar := pre - k
		if cnt, exists := m[tar]; exists {
			ans += cnt
		}
		m[pre] += 1
	}
	return ans
}

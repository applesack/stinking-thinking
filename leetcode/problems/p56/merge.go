package p56

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var ans [][]int
	for _, interval := range intervals {
		length := len(ans)
		if length == 0 || ans[length-1][1] < interval[0] {
			ans = append(ans, interval)
		} else {
			ans[length-1][1] = max(ans[length-1][1], interval[1])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p435

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// 按照每个区间的右边界位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	offset, right := 1, intervals[0][1]
	for _, item := range intervals[1:] {
		if item[0] >= right {
			right = item[1]
			offset++
		}
	}
	// offset: 每次有一个和当前区间a不重叠的区间b, 就加一
	// 用输入区间列表的长度减去不重叠区间的数量, 得到需要移除的区间数量
	return len(intervals) - offset
}

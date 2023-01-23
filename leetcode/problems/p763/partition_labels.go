package p763

import "sort"

// ababcbacadefegdehijhklij
// 012345678901234567890123
// a-------a
//
//	b---b
//	   c--c
//	        d----d
//	           e--e
//	          f
//	            g
//	               h--h
//	                i----i
//	                 j----j
//	                   k
//
// 9,7,8
func partitionLabels(s string) []int {
	m := make(map[int][]int)
	for i, ch := range s {
		if record, ok := m[int(ch)]; ok {
			record[1] = i
		} else {
			record = make([]int, 2)
			record[0], record[1] = i, i
			m[int(ch)] = record
		}
	}

	intervals := reduce(sortValues(m))
	ans := make([]int, len(intervals))
	for i, interval := range intervals {
		ans[i] = interval[1] - interval[0] + 1
	}
	return ans
}

func sortValues(m map[int][]int) [][]int {
	result, i := make([][]int, len(m)), 0
	for _, v := range m {
		result[i] = v
		i++
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}

func reduce(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return make([][]int, 0)
	}

	var reduced [][]int
	reduced = append(reduced, intervals[0])
	pos := 0
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		if interval[0] < reduced[pos][1] {
			reduced[pos][1] = max(reduced[pos][1], interval[1])
		} else {
			reduced = append(reduced, interval)
			pos++
		}
	}

	return reduced
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package p763

func partitionLabelsV2(s string) (partition []int) {
	lastPos := [26]int{}
	// 标记每个字符最后一次出现的位置
	for i, c := range s {
		lastPos[c-'a'] = i
	}
	// 维护两个变量, 开始位置和结束位置
	start, end := 0, 0
	for i, c := range s {
		// 如果当前符号的结束位置比之前记录的结束位置大, 那么更新结束位置
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		}
		// 如果当前符号就是前面标记的结束位置, 那么收集一个结果, 同时更新下一个开始位置
		if i == end {
			partition = append(partition, end-start+1)
			start = end + 1
		}
	}
	return
}

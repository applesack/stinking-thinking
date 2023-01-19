package p119

// 1         0   1
// 1 1       1   1 1
// 1 2 1     2   1 2 1
// 1 3 3 1   3   1 3 3 1
func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	for i := 0; i < len(row); i++ {
		row[i] = 1
		for j := i - 1; j >= 1; j-- {
			row[j] = row[j] + row[j-1]
		}
	}
	return row
}

package p240

func searchMatrix(matrix [][]int, target int) bool {
	ROW, COL := len(matrix), len(matrix[0])
	if target < matrix[0][0] || target > matrix[ROW-1][COL-1] {
		return false
	}
	for r := 0; r < ROW; r++ {
		if matrix[r][0] > target {
			break
		}
		if target >= matrix[r][0] && target <= matrix[r][COL-1] {
			if binarySearch(matrix[r], target) {
				return true
			}
		}
	}
	return false
}

func binarySearch(mat []int, t int) bool {
	l, r, m := 0, len(mat)-1, 0
	for l <= r {
		m = l + (r-l)/2
		if mat[m] == t {
			return true
		} else if mat[m] > t {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return false
}

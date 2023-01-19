package p59

func generateMatrix(n int) [][]int {
	Direction := [][]int{
		{0, 1},  // 右
		{1, 0},  // 下
		{0, -1}, // 左
		{-1, 0}, // 上
	}
	TOP, BOTTOM, LEFT, RIGHT := 0, n-1, 0, n-1
	mat := make([][]int, n)
	for i := 0; i < len(mat); i++ {
		mat[i] = make([]int, n)
	}
	r, c, pos, size, move, direct := 0, 0, 1, n*n, 0, false
	for ; pos <= size; pos++ {
		mat[r][c] = pos
		r, c = r+Direction[move][0], c+Direction[move][1]
		if move == 0 && c == RIGHT {
			TOP, direct = TOP+1, true
		} else if move == 1 && r == BOTTOM {
			RIGHT, direct = RIGHT-1, true
		} else if move == 2 && c == LEFT {
			BOTTOM, direct = BOTTOM-1, true
		} else if move == 3 && r == TOP {
			LEFT, direct = LEFT+1, true
		}
		if direct {
			move = (move + 1) % 4
			direct = false
		}
	}
	return mat
}

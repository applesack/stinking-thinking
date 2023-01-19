package p48

func rotate(matrix [][]int) {
	for r, c := 0, 0; r < len(matrix)/2; r, c = r+1, c+1 {
		rotateMatrix(matrix, r, c, len(matrix)-r*2)
	}
}

func rotateMatrix(mat [][]int, r, c, length int) {
	section := length - 1
	for i := 0; i < section; i++ {
		tl, tr, bl, br := mat[r][c+i], mat[r+i][r+section], mat[r+section-i][c], mat[r+section][c+section-i]
		mat[r][c+i], mat[r+i][r+section], mat[r+section-i][c], mat[r+section][c+section-i] = bl, tl, br, tr
	}
}

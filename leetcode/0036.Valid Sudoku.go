package leetcode

import "strconv"

func isValidSudoku(board [][]byte) bool {
	// convert to board integer
	var boardInt [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			str := string(board[i][j])
			if str == "." {
				boardInt[i][j] = 0
			} else {
				num, _ := strconv.Atoi(str)
				boardInt[i][j] = num
			}
		}
	}

	var tmpHorizantal [10]bool
	var tmpVertical [10]bool
	var tmpSquare [10]bool
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// horizontal scanning
			val := boardInt[i][j]
			if tmpHorizantal[val] {
				return false
			} else if val != 0 {
				tmpHorizantal[val] = true
			}

			// vertical scanning
			val = boardInt[j][i]
			if tmpVertical[val] {
				return false
			} else if val != 0 {
				tmpVertical[val] = true
			}

			// square scanning
			row := (i / 3 * 3) + (j / 3)
			col := ((i % 3) * 3) + j%3
			val = boardInt[row][col]
			if tmpSquare[val] {
				return false
			} else if val != 0 {
				tmpSquare[val] = true
			}
		}
		tmpHorizantal = [10]bool{}
		tmpVertical = [10]bool{}
		tmpSquare = [10]bool{}
	}
	return true
}

package sudokusolution

/*
Sudoku Background
Sudoku is a game played on a 9x9 grid.
The goal of the game is to fill all cells of the grid with digits from 1 to 9,
so that each column, each row, and each of the nine 3x3 sub-grids (also known as blocks) contain all of the digits from 1 to 9.
(More info at: http://en.wikipedia.org/wiki/Sudoku)
Sudoku Solution Validator
Write a function validSolution/ValidateSolution/valid_solution() that accepts a 2D array representing a Sudoku board,
and returns true if it is a valid solution, or false otherwise.
The cells of the sudoku board may also contain 0's, which will represent empty cells.
Boards containing one or more zeroes are considered to be invalid solutions.
The board is always 9 cells by 9 cells, and every cell only contains integers from 0 to 9.
*/

func ValidateSolution(m [][]int) bool {
	// Checking rows and cols
	for r := 0; r < 9; r++ {
		row := make([]int, 9)
		col := make([]int, 9)
		for c := 0; c < 9; c++ {
			if m[r][c] == 0 || m[c][r] == 0 {
				return false
			}
			if row[m[r][c]-1] == 1 || col[m[c][r]-1] == 1 {
				return false
			}
			row[m[r][c]-1] = 1
			col[m[c][r]-1] = 1
		}
	}
	// Checking squares
	for r := 0; r < 9; r += 3 {
		for c := 0; c < 9; c += 3 {
			row := make([]int, 9)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if row[m[r+i][c+j]-1] == 1 {
						return false
					}
					row[m[r+i][c+j]-1] = 1
				}
			}
		}
	}
	return true
}

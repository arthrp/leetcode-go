package leetcode

// 51 - N-Queens
// Queens attack along a row, a column, or a diagonal. Exactly one queen goes
// in each row, so two queens never share a row. cols[c] is true when column
// c already holds a queen. The two diagonal directions are indexed separately:
// squares on one diagonal share row-col, and squares on the other share
// row+col. row-col runs from 1-n to n-1, so adding n-1 keeps that index
// non-negative. row+col runs from 0 to 2n-2.
// Time is O(n!). Extra memory is O(n), plus the boards returned.

func solveNQueens(n int) [][]string {
	result := make([][]string, 0) // solutions may appear in any order
	queens := make([]int, n)      // queens[r] is the column chosen for row r; one queen per row, so rows never clash
	cols := make([]bool, n)       // cols[c] is true when column c already holds a queen
	diag := make([]bool, 2*n-1)   // squares on one diagonal share row-col, shifted into [0, 2n-2]
	anti := make([]bool, 2*n-1)   // squares on the other diagonal share row+col, from 0 to 2n-2

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n { // every row has a queen; copy the column choices into a board
			board := make([]string, n) // n strings of length n
			for r := range queens {
				line := make([]byte, n)
				for c := range line {
					line[c] = '.' // empty square
				}
				line[queens[r]] = 'Q' // queen
				board[r] = string(line)
			}
			result = append(result, board)
			return
		}

		for c := 0; c < n; c++ { // fill the board from this row downward
			d := row - c + n - 1 // row-col runs from 1-n to n-1; adding n-1 keeps the index non-negative
			a := row + c
			if cols[c] || diag[d] || anti[a] {
				continue // queens attack on a column or a diagonal; skip an occupied one
			}
			cols[c], diag[d], anti[a] = true, true, true // occupy the column and both diagonals
			queens[row] = c
			backtrack(row + 1)
			cols[c], diag[d], anti[a] = false, false, false // clear the marks and try the next column from the rows above
		}
	}

	backtrack(0)
	return result
}

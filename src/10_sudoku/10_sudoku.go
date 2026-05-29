package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	row   int
	col   int
	block int
}

func getBlockIndex(row int, col int) int {
	return (row/3)*3 + col/3
}

func checkPosition(blocks [][]bool, columns [][]bool, rows [][]bool, pos *Position, number int) bool {

	if rows[pos.row][number-1] || columns[pos.col][number-1] || blocks[pos.block][number-1] {
		return false
	}
	return true
}

func solveSudoku(board [][]int, pos *Position, blocks [][]bool, columns [][]bool, rows [][]bool) int {
	if pos.col == 9 {
		pos.row++
		pos.col = 0
	}

	if pos.row == 9 {
		return 1
	}

	if board[pos.row][pos.col] != 0 {
		nextPos := Position{row: pos.row, col: pos.col + 1}
		return solveSudoku(board, &nextPos, blocks, columns, rows)
	}
	totalSolutions := 0

	pos.block = getBlockIndex(pos.row, pos.col)

	for number := 1; number <= 9; number++ {
		if checkPosition(blocks, columns, rows, pos, number) {
			blocks[pos.block][number-1] = true
			columns[pos.col][number-1] = true
			rows[pos.row][number-1] = true
			board[pos.row][pos.col] = number

			nextPos := Position{row: pos.row, col: pos.col + 1}
			totalSolutions += solveSudoku(board, &nextPos, blocks, columns, rows)

			blocks[pos.block][number-1] = false
			columns[pos.col][number-1] = false
			rows[pos.row][number-1] = false
			board[pos.row][pos.col] = 0
		}
	}

	return totalSolutions
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	sudoku := make([][]int, 9)
	block := make([][]bool, 9)
	columns := make([][]bool, 9)
	rows := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		block[i] = make([]bool, 9)
		columns[i] = make([]bool, 9)
		rows[i] = make([]bool, 9)
		sudoku[i] = make([]int, 9)
	}

	numberOfHints := 0
	if _, err := fmt.Fscan(in, &numberOfHints); err != nil {
		return
	}
	numberOfSolutions := 0
	var i, j, k int
	for numb := 0; numb < numberOfHints; numb++ {
		fmt.Fscan(in, &i, &j, &k)
		sudoku[i-1][j-1] = k

		b := getBlockIndex(i-1, j-1)
		rows[i-1][k-1] = true
		columns[j-1][k-1] = true
		block[b][k-1] = true
	}

	var pos Position = Position{row: 0, col: 0, block: 0}

	numberOfSolutions = solveSudoku(sudoku, &pos, block, columns, rows)

	fmt.Fprintln(out, numberOfSolutions)
}

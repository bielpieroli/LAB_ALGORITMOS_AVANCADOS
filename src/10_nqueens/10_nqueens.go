package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func isSafe(tabuleiro [][]bool, row int, col int) bool {
	// Verificar a coluna
	for i := 0; i < len(tabuleiro); i++ {
		if tabuleiro[i][col] {
			return false
		}
	}

	// Verificar a diagonal superior esquerda
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if tabuleiro[i][j] {
			return false
		}
	}

	// Verificar a diagonal superior direita
	for i, j := row, col; i >= 0 && j < len(tabuleiro); i, j = i-1, j+1 {
		if tabuleiro[i][j] {
			return false
		}
	}

	return true
}

func placeQueens(tabuleiro [][]bool, row int, solutions *list.List) {

	var lenTab int = len(tabuleiro)

	if row == lenTab {
		answer := list.New()
		for i := 0; i < lenTab; i++ {
			for j := 0; j < lenTab; j++ {
				if tabuleiro[i][j] {
					answer.PushBack(j)
				}
			}
		}
		solutions.PushBack(answer)
	}

	// Possível otimização é colocar o item diretamente na pilha quando encontrar um candidato e desempilhar caso não dê
	for j := 0; j < lenTab; j++ {
		if isSafe(tabuleiro, row, j) {
			tabuleiro[row][j] = true

			placeQueens(tabuleiro, row+1, solutions)

			tabuleiro[row][j] = false
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	solutions := list.New()
	var numQueens int = 1

	for {
		solutions.Init()
		if _, err := fmt.Fscan(in, &numQueens); err != nil {
			return
		}
		if numQueens == 0 {
			break
		}
		tabuleiro := make([][]bool, numQueens)
		for i := 0; i < numQueens; i++ {
			tabuleiro[i] = make([]bool, numQueens)
		}

		placeQueens(tabuleiro, 0, solutions)
		var lenSols int = solutions.Len()
		fmt.Fprintf(out, "[%d,%d]\n", numQueens, lenSols)
		if lenSols == 0 {
			fmt.Fprintln(out, "sem solucao")
		} else {
			primSol := solutions.Front().Value.(*list.List)
			for c := primSol.Front(); c != nil; c = c.Next() {
				fmt.Fprintf(out, "%d ", c.Value.(int)+1)
			}
			fmt.Fprint(out, "\n")
		}
	}
}

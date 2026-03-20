package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var linhas, colunas int
	fmt.Fscan(in, &linhas, &colunas)

	pacmanGrid := make([][]int, linhas)
	for i := 0; i < linhas; i++ {
		pacmanGrid[i] = make([]int, colunas)
	}

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			fmt.Fscan(in, &pacmanGrid[i][j])
		}
	}

	maxPastilhas := calcularCaminho(pacmanGrid, linhas, colunas)
	fmt.Fprintln(out, maxPastilhas)
}

func calcularCaminho(pacmanGrid [][]int, linhas, colunas int) int {
	dpPastilhas := make([][]int, linhas)

	for i := 0; i < linhas; i++ {
		dpPastilhas[i] = make([]int, colunas)
		for j := 0; j < colunas; j++ {
			dpPastilhas[i][j] = -1
		}
	}

	if pacmanGrid[0][0] != -1 {
		dpPastilhas[0][0] = pacmanGrid[0][0]

		for i := 0; i < linhas; i++ {
			for j := 0; j < colunas; j++ {
				if pacmanGrid[i][j] == -1 {
					continue
				}

				if i == 0 && j == 0 {
					continue
				}

				best := -1
				if j > 0 && dpPastilhas[i][j-1] >= 0 {
					best = dpPastilhas[i][j-1]
				}
				if i > 0 && dpPastilhas[i-1][j] >= 0 && dpPastilhas[i-1][j] > best {
					best = dpPastilhas[i-1][j]
				}
				if best >= 0 {
					dpPastilhas[i][j] = best + pacmanGrid[i][j]
				}
			}
		}
	}

	if dpPastilhas[linhas-1][colunas-1] < 0 {
		return -1
	}
	return dpPastilhas[linhas-1][colunas-1]
}

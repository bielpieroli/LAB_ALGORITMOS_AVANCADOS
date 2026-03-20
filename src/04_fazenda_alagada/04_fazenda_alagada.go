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

	patoGrid := make([][]int, linhas)
	for i := 0; i < linhas; i++ {
		patoGrid[i] = make([]int, colunas)
	}

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			fmt.Fscan(in, &patoGrid[i][j])
		}
	}
	// fmt.Println(patoGrid)

	maxGraos := natacaoPato(patoGrid, linhas, colunas)

	fmt.Fprintln(out, maxGraos)
}

func natacaoPato(patoGrid [][]int, linhas, colunas int) int {
	dpGraos := make([][]int, linhas)

	vetorIda := make([]int, colunas)
	vetorVolta := make([]int, colunas)

	for i := 0; i < linhas; i++ {
		dpGraos[i] = make([]int, colunas)
		for j := 0; j < colunas; j++ {
			dpGraos[i][j] = -1
		}
	}

	for j := 0; j < colunas; j++ {
		vetorIda[j] = -1
		vetorVolta[j] = -1
	}

	if patoGrid[0][0] != -1 || patoGrid[linhas-1][colunas-1] != -1 {
		dpGraos[0][0] = patoGrid[0][0]

		for i := 0; i < linhas; i++ {
			vetorIda[0] = dpGraos[i][0]
			for j := 0; j < colunas; j++ {
				if patoGrid[i][j] == -1 {
					vetorIda[j] = -1
					continue
				}

				if i == 0 && j == 0 {
					continue
				}

				if j > 0 && vetorIda[j-1] >= 0 {
					vetorIda[j] = vetorIda[j-1] + patoGrid[i][j]
				}
				if i > 0 && dpGraos[i-1][j] >= 0 && dpGraos[i-1][j]+patoGrid[i][j] > vetorIda[j] {
					vetorIda[j] = dpGraos[i-1][j] + patoGrid[i][j]
				}
			}
			// fmt.Println("ida:", vetorIda)
			vetorVolta[colunas-1] = dpGraos[i][colunas-1]
			if i != 0 {
				for j := colunas - 1; j >= 0; j-- {
					if patoGrid[i][j] == -1 {
						vetorVolta[j] = -1
						continue
					}
					if j < colunas-1 && vetorVolta[j+1] >= 0 {
						vetorVolta[j] = vetorVolta[j+1] + patoGrid[i][j]
					}
					if i > 0 && dpGraos[i-1][j] >= 0 && dpGraos[i-1][j]+patoGrid[i][j] > vetorVolta[j] {
						vetorVolta[j] = dpGraos[i-1][j] + patoGrid[i][j]
					}
					dpGraos[i][j] = max(vetorVolta[j], vetorIda[j])
				}
				// fmt.Println("volta:", vetorVolta)
			} else {
				for j := 0; j < colunas; j++ {
					dpGraos[i][j] = max(vetorVolta[j], vetorIda[j])
				}
			}
			// fmt.Println(dpGraos)
		}
	}

	if dpGraos[linhas-1][colunas-1] < 0 {
		return -1
	}
	return dpGraos[linhas-1][colunas-1]
}

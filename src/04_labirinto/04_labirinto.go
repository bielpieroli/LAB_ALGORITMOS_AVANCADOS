package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const modulo = 1000000007

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var grau int
	fmt.Fscan(in, &grau)

	labirinto := make([][]string, grau)
	temp := make([]string, grau)

	for i := 0; i < grau; i++ {
		fmt.Fscan(in, &temp[i])
	}

	for i := 0; i < grau; i++ {
		labirinto[i] = strings.Split(temp[i], "")
	}

	qntdCaminhos := calcularCaminho(labirinto, grau)
	fmt.Fprintln(out, qntdCaminhos)
}

func calcularCaminho(labirinto [][]string, grau int) int {

	calculoCaminhos := make([][]int, grau)

	for i := 0; i < grau; i++ {
		calculoCaminhos[i] = make([]int, grau)
	}

	calculoCaminhos[0][0] = 1

	for i := 0; i < grau; i++ {
		for j := 0; j < grau; j++ {
			if labirinto[i][j] == "*" {
				calculoCaminhos[i][j] = 0
			} else {
				if i > 0 {
					calculoCaminhos[i][j] = (calculoCaminhos[i][j] + calculoCaminhos[i-1][j]) % modulo
				}
				if j > 0 {
					calculoCaminhos[i][j] = (calculoCaminhos[i][j] + calculoCaminhos[i][j-1]) % modulo
				}
			}
		}
	}

	return calculoCaminhos[grau-1][grau-1] % modulo
}

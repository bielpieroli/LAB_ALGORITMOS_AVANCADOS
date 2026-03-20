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

	var qntdTestes int
	fmt.Fscan(in, &qntdTestes)

	// fmt.Fprint(out, qntdTestes, "\n")

	for i := 0; i < qntdTestes; i++ {
		var linhas, colunas int
		fmt.Fscan(in, &linhas, &colunas)

		// fmt.Fprint(out, linhas, colunas, "\n")

		gridPintura := make([][]string, linhas)
		temp := make([]string, linhas)

		for i := 0; i < linhas; i++ {
			fmt.Fscan(in, &temp[i])
		}

		for i := 0; i < linhas; i++ {
			gridPintura[i] = strings.Split(temp[i], "")
		}

		menorAltura, menorLargura := calcularMenorRetangulo(gridPintura, linhas, colunas)
		// fmt.Fprintln(out, gridPintura)
		fmt.Fprintln(out, menorAltura, menorLargura)
	}
}

func calcularMenorRetangulo(gridPintura [][]string, linhas int, colunas int) (int, int) {
	menorAltura := 0
	menorLargura := 0
	minJIndex := colunas
	maxJIndex := 0

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			if gridPintura[i][j] == "B" {
				if menorAltura == 0 {
					menorAltura = linhas - i
				}
				if j < minJIndex {
					minJIndex = j
				}
				if j > maxJIndex || maxJIndex < minJIndex {
					maxJIndex = j
				}
			}
		}
	}
	// fmt.Println("min:", minJIndex, "max:", maxJIndex)
	if maxJIndex < minJIndex {
		maxJIndex = minJIndex - 1
	}
	menorLargura = maxJIndex - minJIndex + 1

	return menorAltura, menorLargura
}

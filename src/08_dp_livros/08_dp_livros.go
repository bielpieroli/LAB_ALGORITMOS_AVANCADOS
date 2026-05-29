package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1e9 + 7

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var numLivros int
	if _, err := fmt.Fscan(in, &numLivros); err != nil {
		return
	}

	livros := make([]int, numLivros)
	for i := 0; i < numLivros; i++ {
		fmt.Fscan(in, &livros[i])
	}

	dp := make([]int, numLivros)
	totalGeral := 0

	for i := 0; i < numLivros; i++ {
		// O livro é uma subsequência por si só, então, comecamos com 1
		dp[i] = 1

		// j < i para ver apenas os livros anteriores
		for j := 0; j < i; j++ {
			// Se o livro anterior for menor que o atual, podemos estender as sequências
			if livros[j] < livros[i] {
				dp[i] = (dp[i] + dp[j]) % MOD
			}
		}

		totalGeral = (totalGeral + dp[i]) % MOD
	}

	fmt.Fprintln(out, totalGeral)
}

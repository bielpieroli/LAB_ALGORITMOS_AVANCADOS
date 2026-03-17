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

	var numLivros, orcamentoMax int
	fmt.Fscan(in, &numLivros, &orcamentoMax)

	precos := make([]int, numLivros)
	pags := make([]int, numLivros)

	for i := 0; i < numLivros; i++ {
		fmt.Fscan(in, &precos[i])
	}

	for i := 0; i < numLivros; i++ {
		fmt.Fscan(in, &pags[i])
	}

	dp := make([]int, orcamentoMax+1)

	for i := 0; i < numLivros; i++ {
		for j := orcamentoMax; j >= precos[i]; j-- {
			valor := dp[j-precos[i]] + pags[i]
			if valor > dp[j] {
				dp[j] = valor
			}
		}
	}
	// fmt.Fprint(out, "NumLivros:", numLivros, "\n")
	// fmt.Fprint(out, "OrcamentoMax:", orcamentoMax, "\n")
	// fmt.Fprint(out, "Precos:", precos, "\n")
	// fmt.Fprint(out, "Pags:", pags, "\n")
	// fmt.Fprint(out, "DP:", dp, "\n")
	fmt.Fprintln(out, dp[orcamentoMax])
}

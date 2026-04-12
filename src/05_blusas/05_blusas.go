package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1000000007

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var qntdConvidados int
	var qntdMaxBlusas int

	fmt.Fscan(in, &qntdConvidados, &qntdMaxBlusas)

	blusaParaPessoas := make([][]int, qntdMaxBlusas+1)

	var qntdBlusas int

	for pessoa := 0; pessoa < qntdConvidados; pessoa++ {
		fmt.Fscan(in, &qntdBlusas)
		for blusa := 0; blusa < qntdBlusas; blusa++ {
			var blusaEntrada int
			fmt.Fscan(in, &blusaEntrada)
			blusaParaPessoas[blusaEntrada] = append(blusaParaPessoas[blusaEntrada], pessoa)
		}
	}

	dp := make([]int64, 1<<qntdConvidados)
	dp[0] = 1

	for blusa := 1; blusa <= qntdMaxBlusas; blusa++ {

		newDP := make([]int64, 1<<qntdConvidados)
		copy(newDP, dp)

		for mask := 0; mask < (1 << qntdConvidados); mask++ {

			if dp[mask] == 0 { // Se não há maneiras de vestir as pessoas com as blusas anteriores
				continue
			}

			for _, pessoa := range blusaParaPessoas[blusa] {

				if (mask & (1 << pessoa)) != 0 { // Se a pessoa já tem uma blusa, não pode usar esta blusa
					continue
				}

				novoMask := mask | (1 << pessoa) // Adiciona a pessoa ao conjunto
				newDP[novoMask] = (newDP[novoMask] + dp[mask]) % MOD
			}
		}

		dp = newDP
	}

	fullMask := (1 << qntdConvidados) - 1
	fmt.Fprintln(out, dp[fullMask])
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var numTestes int
	fmt.Fscan(in, &numTestes)

	for i := 0; i < numTestes; i++ {
		var valorDinheiro, qntdMoedas int
		fmt.Fscan(in, &valorDinheiro, &qntdMoedas)

		moedas := make([]int, qntdMoedas)
		for j := 0; j < qntdMoedas; j++ {
			fmt.Fscan(in, &moedas[j])
		}

		dp := make([]int, valorDinheiro+1)
		for j := 0; j <= valorDinheiro; j++ {
			if contains(moedas, j) {
				dp[j] = 1
			} else {
				dp[j] = math.MaxInt32
			}
		}

		for i := 0; i < valorDinheiro+1; i++ {
			for _, moeda := range moedas {
				if i-moeda >= 0 {
					dp[i] = min(dp[i], dp[i-moeda]+1)
				}
			}
		}
		fmt.Println(dp[valorDinheiro])
	}
}

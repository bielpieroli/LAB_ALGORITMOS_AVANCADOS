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

	var distancia int
	fmt.Fscan(in, &distancia)

	saltos := make([]int, distancia)

	for i := 0; i < distancia; i++ {
		fmt.Fscan(in, &saltos[i])
	}

	dp := make([]int, distancia)

	for i := 0; i < distancia && dp[distancia-1] == 0; i++ {
		if dp[i] == 0 && i != 0 {
			continue
		}
		for j := 1; j <= saltos[i] && dp[distancia-1] == 0; j++ {
			if dp[i+j] != 0 {
				dp[i+j] = min(dp[i]+1, dp[i+j])
			} else {
				dp[i+j] = dp[i] + 1
			}

		}
	}
	// fmt.Fprint(out, "DP:", dp, "\n")
	// fmt.Fprint(out, "Saltos:", saltos, "\n")
	if dp[distancia-1] != 0 {
		fmt.Fprintln(out, dp[distancia-1])
	} else {
		fmt.Fprintln(out, "Salto impossivel")
	}
}

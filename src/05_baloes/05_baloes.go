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
	var qntdBaloes int
	fmt.Fscan(in, &qntdBaloes)

	baloes := make([]int, qntdBaloes+2)
	baloes[0] = 1
	baloes[qntdBaloes+2-1] = 1

	for i := 1; i <= qntdBaloes; i++ {
		fmt.Fscan(in, &baloes[i])
	}

	dp := make([][]int, qntdBaloes+2)
	for i := 0; i < qntdBaloes+2; i++ {
		dp[i] = make([]int, qntdBaloes+2)
	}

	// fmt.Fprintln(out, baloes)

	for i := 1; i <= qntdBaloes; i++ {
		for j := 1; j <= qntdBaloes-i+1; j++ {
			r := j + i - 1
			for k := j; k <= r; k++ {
				dp[j][r] = max(
					dp[j][r],
					dp[j][k-1]+dp[k+1][r]+baloes[j-1]*baloes[k]*baloes[r+1],
				)
			}
		}
	}

	// fmt.Fprintln(out, dp)
	fmt.Fprint(out, dp[1][qntdBaloes])

}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var qntdGomos int
	if _, err := fmt.Fscan(in, &qntdGomos); err != nil {
		return
	}

	gomos := make([]int, qntdGomos)
	for i := 0; i < qntdGomos; i++ {
		if _, err := fmt.Fscan(in, &gomos[i]); err != nil {
			return
		}
	}

	dp := make([]int, qntdGomos)
	for i := 0; i < qntdGomos; i++ {
		dp[i] = gomos[i]
	}

	for i := 0; i < qntdGomos; i++ {
		for j := i; j < qntdGomos; j++ {
			if i+j+2 > qntdGomos {
				break
			}
			if dp[i]+dp[j] > dp[i+j+1] {
				dp[i+j+1] = dp[i] + dp[j]
			}
		}
	}

	fmt.Println(dp[qntdGomos-1])
}

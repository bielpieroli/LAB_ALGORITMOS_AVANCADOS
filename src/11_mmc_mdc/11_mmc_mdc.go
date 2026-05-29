package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func mmc(a int, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return abs(a*b) / mdc(a, b)
}

func mdc(a int, b int) int {
	if b == 0 {
		return a
	}
	return mdc(b, a%b)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var numOfTests int
	if _, err := fmt.Fscan(in, &numOfTests); err != nil {
		return
	}

	var a, b int

	for i := 0; i < numOfTests; i++ {
		if _, err := fmt.Fscan(in, &a, &b); err != nil {
			return
		}
		fmt.Fprintln(out, mdc(a, b), mmc(a, b))
	}
}

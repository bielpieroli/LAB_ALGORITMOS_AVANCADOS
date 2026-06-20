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

	var n int

	if _, err := fmt.Fscan(in, &n); err != nil {
		return
	}

	if n%2 == 0 {
		fmt.Fprintln(out, "Primeiro")
	} else {
		fmt.Fprintln(out, "Segundo")
	}
}

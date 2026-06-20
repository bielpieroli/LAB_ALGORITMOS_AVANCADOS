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

	var num_torres int
	if _, err := fmt.Fscan(in, &num_torres); err != nil {
		return
	}

	var xor int

	var torres int
	for i := 0; i < num_torres; i++ {
		if _, err := fmt.Fscan(in, &torres); err != nil {
			return
		}
		xor = xor ^ torres
	}

	if xor == 0 {
		fmt.Fprintln(out, "Segundo")
	} else {
		fmt.Fprintln(out, "Primeiro")
	}
}

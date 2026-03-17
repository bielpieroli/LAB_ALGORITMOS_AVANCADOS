package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	if _, err := fmt.Fscan(in, &t); err != nil {
		return
	}

	for i := 0; i < t; i++ {
		var x, y int
		if _, err := fmt.Fscan(in, &x, &y); err != nil {
			return
		}

		// Regras:
		// - Se x for ímpar -> não é possível dividir igualmente -> NO
		// - Se x == 0 e y for ímpar -> não é possível dividir igualmente -> NO
		if (x == 0 && y%2 == 1) || (x%2 == 1) {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}

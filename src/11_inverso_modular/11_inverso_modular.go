package main

import (
	"bufio"
	"fmt"
	"os"
)

// x é tal que a*x + m*y = 1, logo a*x ≡ 1 (mod m)

func euclides_estendido(a int, b int) (int, int, int) {
	if b == 0 {
		// Caso base: mdc(a,0) = a, a*1 + 0*0 = a
		return a, 1, 0
	}

	mdc, x1, y1 := euclides_estendido(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return mdc, x, y
}

func inverso_modular(a, m int) int {
	mdc, x, _ := euclides_estendido(a, m)
	if mdc != 1 {
		return -1
	}
	println("x:", x)
	if x < 0 {
		x = (x%m + m) % m
	}
	return x % m
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
		fmt.Fprintln(out, inverso_modular(a, b))
	}
}

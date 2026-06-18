package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Ponto struct {
	x, y int
}

type Retangulo struct {
	verticeMin, verticeMax Ponto
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var retangulo1 Retangulo = Retangulo{
		verticeMin: Ponto{
			x: math.MaxInt,
			y: math.MaxInt,
		},
		verticeMax: Ponto{
			x: math.MinInt,
			y: math.MinInt,
		},
	}
	var retangulo2 Retangulo = Retangulo{
		verticeMin: Ponto{
			x: math.MaxInt,
			y: math.MaxInt,
		},
		verticeMax: Ponto{
			x: math.MinInt,
			y: math.MinInt,
		},
	}
	var x, y int

	for i := 0; i < 4; i++ {
		if _, err := fmt.Fscan(in, &x, &y); err != nil {
			return
		}
		if x < retangulo1.verticeMin.x {
			retangulo1.verticeMin.x = x
		}
		if y < retangulo1.verticeMin.y {
			retangulo1.verticeMin.y = y
		}
		if x > retangulo1.verticeMax.x {
			retangulo1.verticeMax.x = x
		}
		if y > retangulo1.verticeMax.y {
			retangulo1.verticeMax.y = y
		}
	}

	for i := 0; i < 4; i++ {
		if _, err := fmt.Fscan(in, &x, &y); err != nil {
			return
		}
		if x < retangulo2.verticeMin.x {
			retangulo2.verticeMin.x = x
		}
		if y < retangulo2.verticeMin.y {
			retangulo2.verticeMin.y = y
		}
		if x > retangulo2.verticeMax.x {
			retangulo2.verticeMax.x = x
		}
		if y > retangulo2.verticeMax.y {
			retangulo2.verticeMax.y = y
		}
	}

	// fmt.Fprintln(out, retangulo1.verticeMin.x, retangulo1.verticeMin.y, retangulo1.verticeMax.x, retangulo1.verticeMax.y)

	// fmt.Fprintln(out, retangulo2.verticeMin.x, retangulo2.verticeMin.y, retangulo2.verticeMax.x, retangulo2.verticeMax.y)

	if (retangulo1.verticeMin.x <= retangulo2.verticeMax.x && retangulo1.verticeMax.x >= retangulo2.verticeMin.x) &&
		(retangulo1.verticeMin.y <= retangulo2.verticeMax.y && retangulo1.verticeMax.y >= retangulo2.verticeMin.y) {
		fmt.Fprint(out, "SIM\n")
	} else {
		fmt.Fprint(out, "NAO\n")
	}

}

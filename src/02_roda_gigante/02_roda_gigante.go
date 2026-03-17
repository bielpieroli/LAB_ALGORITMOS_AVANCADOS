package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var numOfTestes int
	if _, err := fmt.Fscan(in, &numOfTestes); err != nil {
		return
	}

	var numCriancas, pesoMaximo int

	for i := 0; i < numOfTestes; i++ {
		if _, err := fmt.Fscan(in, &numCriancas, &pesoMaximo); err != nil {
			return
		}

		criancas := make([]int, numCriancas)
		for i := 0; i < numCriancas; i++ {
			if _, err := fmt.Fscan(in, &criancas[i]); err != nil {
				return
			}
		}

		sort.Sort(sort.Reverse(sort.IntSlice(criancas)))

		var numCabines int

		for len(criancas) > 0 {
			if len(criancas) == 1 {
				criancas = criancas[1:]
			} else {
				pesoMaisPesado := criancas[0]
				pesoMaisLeve := criancas[len(criancas)-1]
				if pesoMaisPesado+pesoMaisLeve <= pesoMaximo {
					criancas = criancas[:len(criancas)-1]
				}
				criancas = criancas[1:]
			}
			numCabines++
		}

		fmt.Println(numCabines)
	}
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Sprinkler struct {
	comp, raio          int
	cobertura_intervalo [2]float64
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var qntd_sprinklers, comprimento, largura int

	if _, err := fmt.Fscan(in, &qntd_sprinklers, &comprimento, &largura); err != nil {
		return
	}
	sprinkler := make([]Sprinkler, 0, qntd_sprinklers)

	var circ_comp, circ_raio int
	for i := 0; i < qntd_sprinklers; i++ {
		if _, err := fmt.Fscan(in, &circ_comp, &circ_raio); err != nil {
			return
		}

		metade_largura := float64(largura) / 2.0

		if float64(circ_raio) >= metade_largura {
			var s Sprinkler
			s.comp = circ_comp
			s.raio = circ_raio

			cobertura := math.Sqrt(float64(s.raio*s.raio) - (metade_largura * metade_largura))

			s.cobertura_intervalo[0] = float64(s.comp) - cobertura
			s.cobertura_intervalo[1] = float64(s.comp) + cobertura

			sprinkler = append(sprinkler, s)
		} else {
			qntd_sprinklers--
		}
	}

	if len(sprinkler) == 0 {
		fmt.Fprintln(out, -1)
		return
	}

	sort.Slice(sprinkler, func(i, j int) bool {
		if sprinkler[i].cobertura_intervalo[0] != sprinkler[j].cobertura_intervalo[0] {
			return sprinkler[i].cobertura_intervalo[0] < sprinkler[j].cobertura_intervalo[0]
		}

		// Se empatar pelo início, ordena pelo fim do intervalo de cobertura (maior primeiro)
		return sprinkler[i].cobertura_intervalo[1] > sprinkler[j].cobertura_intervalo[1]
	})

	// fmt.Fprintln(out, "largura:", largura, "comprimento:", comprimento)
	// fmt.Fprintln(out, sprinkler)

	if sprinkler[0].cobertura_intervalo[0] > 0 {
		fmt.Fprintln(out, -1)
		return
	}

	qntd_necessaria := 0
	atual_coberto := 0.0
	i := 0

	for atual_coberto < float64(comprimento) {
		melhor_fim := atual_coberto
		valido := false

		for i < len(sprinkler) && sprinkler[i].cobertura_intervalo[0] <= atual_coberto {
			valido = true

			if sprinkler[i].cobertura_intervalo[1] > melhor_fim {
				melhor_fim = sprinkler[i].cobertura_intervalo[1]
			}
			i++
		}

		if !valido || melhor_fim == atual_coberto {
			fmt.Fprintln(out, -1)
			return
		}

		atual_coberto = melhor_fim
		qntd_necessaria++
	}

	fmt.Fprintln(out, qntd_necessaria)
}

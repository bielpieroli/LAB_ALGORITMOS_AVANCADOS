package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var numOfTestes int
	if _, err := fmt.Fscan(in, &numOfTestes); err != nil {
		return
	}

	for i := 0; i < numOfTestes; i++ {
		var numEntregas, numMotoboys int
		if _, err := fmt.Fscan(in, &numEntregas, &numMotoboys); err != nil {
			return
		}

		entregas := make([][2]int, numEntregas)
		for j := 0; j < numEntregas; j++ {
			var tempInicio, tempFim int
			if _, err := fmt.Fscan(in, &tempInicio, &tempFim); err != nil {
				return
			}
			entregas[j] = [2]int{tempInicio, tempFim}
		}

		// Ordena as entregas pelo horário de início e depois pelo horário de término
		// Isso garante que as entregas sejam processadas na ordem correta
		// para maximizar o número de entregas que podem ser atribuídas aos motoboys
		for j := 0; j < numEntregas; j++ {
			for k := j + 1; k < numEntregas; k++ {
				if entregas[j][0] > entregas[k][0] || (entregas[j][0] == entregas[k][0] && entregas[j][1] > entregas[k][1]) {
					entregas[j], entregas[k] = entregas[k], entregas[j]
				}
			}
		}

		motoBoyFinishTimes := make([]int, numMotoboys)
		canComplete := make([]bool, numEntregas)
		for j := 0; j < numEntregas; j++ {
			entregaInicio, entregaFim := entregas[j][0], entregas[j][1]
			assigned := false
			for k := 0; k < numMotoboys; k++ {
				if motoBoyFinishTimes[k] <= entregaInicio {
					motoBoyFinishTimes[k] = entregaFim
					assigned = true
					break
				}
			}
			canComplete[j] = assigned
		}

		count := 0
		for j := 0; j < numEntregas; j++ {
			if canComplete[j] {
				count++
			}
		}
		fmt.Println(count)
	}
}

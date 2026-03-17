package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var numTarefas int
	fmt.Fscan(in, &numTarefas)
	tarefas := make([][2]int, numTarefas)

	// tarefas[0][0] = duracao
	// tarefas[0][1] = prazo
	for i := 0; i < numTarefas; i++ {
		fmt.Fscan(in, &tarefas[i][0], &tarefas[i][1])
	}

	sort.Slice(tarefas, func(i, j int) bool {
		if tarefas[i][0] == tarefas[j][0] {
			return tarefas[i][1] < tarefas[j][1]
		}
		return tarefas[i][0] < tarefas[j][0]
	})

	tarefasFinishTime := 0
	pontos := 0

	for _, tarefa := range tarefas {
		tarefasFinishTime += tarefa[0]
		pontos += tarefa[1] - tarefasFinishTime
	}

	fmt.Fprintln(out, pontos)
}

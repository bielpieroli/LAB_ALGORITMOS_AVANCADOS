package main

import (
	"bufio"
	"fmt"
	"os"
)

func DFS(adj [][]int, start int, visited []bool) int {
	stack := make([]int, 0, len(adj))
	stack = append(stack, start)
	visited[start] = true
	cnt := 1
	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, u := range adj[v] {
			if !visited[u] {
				visited[u] = true
				cnt++
				stack = append(stack, u)
			}
		}
	}
	return cnt
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var numCidades, numVoos int
	if _, err := fmt.Fscan(in, &numCidades, &numVoos); err != nil {
		return
	}

	listaAdjIda := make([][]int, numCidades)
	listaAdjVolta := make([][]int, numCidades)
	for i := 0; i < numVoos; i++ {
		var origem, destino int
		fmt.Fscan(in, &origem, &destino)
		origem--
		destino--
		if origem >= 0 && origem < numCidades && destino >= 0 && destino < numCidades {
			listaAdjIda[origem] = append(listaAdjIda[origem], destino)
			listaAdjVolta[destino] = append(listaAdjVolta[destino], origem)
		}
	}

	visitasIda := make([]bool, numCidades)
	totalCidadesIda := DFS(listaAdjIda, 0, visitasIda)
	if totalCidadesIda != numCidades {
		fmt.Fprintln(out, "NAO")
		return
	}

	visitasVolta := make([]bool, numCidades)
	totalCidadesVolta := DFS(listaAdjVolta, 0, visitasVolta)
	if totalCidadesVolta != numCidades {
		fmt.Fprintln(out, "NAO")
		return
	}

	fmt.Fprintln(out, "SIM")

}

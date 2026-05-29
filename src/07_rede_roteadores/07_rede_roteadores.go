package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var visitado []bool
var pilha []int

func dfs1(u int, adj [][]int) {
	visitado[u] = true
	for _, v := range adj[u] {
		if !visitado[v] {
			dfs1(v, adj)
		}
	}
	pilha = append(pilha, u)
}

// Retorna a lista de nós que pertencem a este componente
func dfs2(u int, rev [][]int, componente *[]int) {
	visitado[u] = true
	*componente = append(*componente, u)
	for _, v := range rev[u] {
		if !visitado[v] {
			dfs2(v, rev, componente)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, M int
	if _, err := fmt.Fscan(in, &N, &M); err != nil {
		return
	}

	conexoes := make([][]int, N+1)
	conexoesReversas := make([][]int, N+1)
	for i := 0; i < M; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		conexoes[u] = append(conexoes[u], v)
		conexoesReversas[v] = append(conexoesReversas[v], u)
	}

	visitado = make([]bool, N+1)
	pilha = []int{}

	for i := 1; i <= N; i++ {
		if !visitado[i] {
			dfs1(i, conexoes)
		}
	}

	for i := range visitado {
		visitado[i] = false
	}

	type subRede struct {
		menor int
		nos   []int
	}

	var subRedes []subRede

	for i := len(pilha) - 1; i >= 0; i-- {
		u := pilha[i]
		if !visitado[u] {
			var atual []int
			dfs2(u, conexoesReversas, &atual)
			sort.Ints(atual)
			subRedes = append(subRedes, subRede{menor: atual[0], nos: atual})
		}
	}

	sort.Slice(subRedes, func(i, j int) bool {
		return subRedes[i].menor < subRedes[j].menor
	})

	for _, s := range subRedes {
		fmt.Fprint(out, "[")
		for idx, no := range s.nos {
			if idx > 0 {
				fmt.Fprint(out, ",")
			}
			fmt.Fprint(out, no)
		}
		fmt.Fprintln(out, "]")
	}
}

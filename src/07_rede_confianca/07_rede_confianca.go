package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	visitado []bool
	pilha    []int
)

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

	var numUsuarios, numConexoes int
	for {
		if _, err := fmt.Fscan(in, &numUsuarios, &numConexoes); err != nil || (numUsuarios == 0 && numConexoes == 0) {
			break
		}

		adj := make([][]int, numUsuarios+1)
		rev := make([][]int, numUsuarios+1)

		for i := 0; i < numConexoes; i++ {
			var origem, destino, relacao int
			fmt.Fscan(in, &origem, &destino, &relacao)

			adj[origem] = append(adj[origem], destino)
			rev[destino] = append(rev[destino], origem)

			if relacao == 2 {
				adj[destino] = append(adj[destino], origem)
				rev[origem] = append(rev[origem], destino)
			}
		}

		// DFS preenchendo a pilha a partir do grafo original
		visitado = make([]bool, numUsuarios+1)
		pilha = []int{}
		for i := 1; i <= numUsuarios; i++ {
			if !visitado[i] {
				dfs1(i, adj)
			}
		}

		// DFS no grafo REVERSO seguindo a ordem inversa da pilha
		for i := range visitado {
			visitado[i] = false
		}

		type Componentes struct {
			menorNo int
			tamanho int
		}
		var componentes []Componentes

		for i := len(pilha) - 1; i >= 0; i-- {
			u := pilha[i]
			if !visitado[u] {
				var compAtual []int
				dfs2(u, rev, &compAtual)

				menor := compAtual[0]
				for _, no := range compAtual {
					if no < menor {
						menor = no
					}
				}
				componentes = append(componentes, Componentes{menorNo: menor, tamanho: len(compAtual)})
			}
		}

		if len(componentes) == 1 {
			fmt.Fprintln(out, "confianca total")
		} else {
			sort.Slice(componentes, func(i, j int) bool {
				return componentes[i].menorNo < componentes[j].menorNo
			})

			for i, scc := range componentes {
				fmt.Fprintf(out, "[%d,%d]", i+1, scc.tamanho)
			}
			fmt.Fprintln(out)
		}
	}
}

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const MOD = int64(1e9 + 7)
const INF64 = int64(1e18)

// Estrutura auxiliar para repsentar as arestas do grafo
type Edge struct {
	to     int
	weight int
}

// Estrutura para representar os itens na fila de prioridade
type Item struct {
	node int
	dist int64
}

type PriorityQueue []*Item

func (FilaPrioridade PriorityQueue) Len() int {
	return len(FilaPrioridade)
}

func (FilaPrioridade PriorityQueue) Less(i, j int) bool {
	return FilaPrioridade[i].dist < FilaPrioridade[j].dist
}

func (FilaPrioridade PriorityQueue) Swap(i, j int) {
	FilaPrioridade[i], FilaPrioridade[j] = FilaPrioridade[j], FilaPrioridade[i]
}

func (FilaPrioridade *PriorityQueue) Push(x interface{}) {
	*FilaPrioridade = append(*FilaPrioridade, x.(*Item))
}

func (FilaPrioridade *PriorityQueue) Pop() interface{} {
	old := *FilaPrioridade
	n := len(old)
	item := old[n-1]
	*FilaPrioridade = old[0 : n-1]
	return item
}

func main() {
	// Scanner com buffer maior para suportar entradas grandes
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var numVertices, numArestas int
	fmt.Fscan(in, &numVertices, &numArestas)

	adj := make([][]Edge, numVertices+1)
	for i := 0; i < numArestas; i++ {
		var origem, destino, peso int
		fmt.Fscan(in, &origem, &destino, &peso)
		adj[origem] = append(adj[origem], Edge{to: destino, weight: peso})
	}

	// São criadas estruturas adicionais para serem atualizadas ao longo do Dijkstra
	dist := make([]int64, numVertices+1)
	ways := make([]int64, numVertices+1)
	minEdges := make([]int, numVertices+1)
	maxEdges := make([]int, numVertices+1)

	// Relembrando que o algoritmo de Dijisktra começa com as distâncias infinitas, exceto para o vértice de origem
	for i := 1; i <= numVertices; i++ {
		dist[i] = INF64
		minEdges[i] = 1e9
		maxEdges[i] = -1e9
	}

	dist[1] = 0
	ways[1] = 1
	minEdges[1] = 0
	maxEdges[1] = 0

	// Adiciona o primeiro node a fila de prioridade
	filaPrioridade := &PriorityQueue{}

	// A fila de prioridade é inicializada como um heap
	heap.Init(filaPrioridade)
	heap.Push(filaPrioridade, &Item{node: 1, dist: 0})

	for filaPrioridade.Len() > 0 {
		current := heap.Pop(filaPrioridade).(*Item)
		noAtual := current.node
		d := current.dist

		// Se já encontramos um caminho menor, não precisamos processar este nó
		if d > dist[noAtual] {
			continue
		}

		// Para cada conexão que esse nó faz, verificamos se encontramos um caminho melhor ou igual
		for _, edge := range adj[noAtual] {
			noConectado := edge.to
			newDist := dist[noAtual] + int64(edge.weight)

			// Caso 1 Encontramos um caminho menor
			if newDist < dist[noConectado] {

				dist[noConectado] = newDist

				ways[noConectado] = ways[noAtual]

				minEdges[noConectado] = minEdges[noAtual] + 1
				maxEdges[noConectado] = maxEdges[noAtual] + 1

				heap.Push(filaPrioridade, &Item{node: noConectado, dist: newDist})

				// Caso 2 Encontramos outro caminho com o mesmo custo mínimo
			} else if newDist == dist[noConectado] {

				ways[noConectado] = (ways[noConectado] + ways[noAtual]) % MOD

				minEdges[noConectado] = min(minEdges[noConectado], minEdges[noAtual]+1)
				maxEdges[noConectado] = max(maxEdges[noConectado], maxEdges[noAtual]+1)
			}
		}
	}

	fmt.Fprintf(out, "%d %d %d %d\n", dist[numVertices], ways[numVertices], minEdges[numVertices], maxEdges[numVertices])
}

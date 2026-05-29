package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Ponto struct {
	x, y float64
}

func calcularDistancia(p1, p2 Ponto) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

type Aresta struct {
	u, v int
	dist float64
}

type DisjointUnion struct {
	pai []int
}

func createDisjointUnion(n int) *DisjointUnion {
	pai := make([]int, n+1)
	for i := range pai {
		pai[i] = i
	}
	return &DisjointUnion{pai: pai}
}

func (d *DisjointUnion) Find(i int) int {
	if d.pai[i] == i {
		return i
	}
	d.pai[i] = d.Find(d.pai[i])
	return d.pai[i]
}

// Unem grupos diferentes e retorna true se a união foi feita, ou false se já estavam no mesmo grupo (ou seja, já estavam conectados)
func (d *DisjointUnion) Union(i, j int) bool {
	raizI := d.Find(i)
	raizJ := d.Find(j)
	if raizI != raizJ {
		d.pai[raizI] = raizJ
		return true
	}
	return false
}

func main() {

	// Encontrar via Kruskal a árvore generadora mínima (MST) de um grafo completo, no qual os vértices são os prédios e as arestas são as distâncias entre eles ..

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	if _, err := fmt.Fscan(in, &N); err != nil {
		return
	}

	predios := make([]Ponto, N+1)
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &predios[i].x, &predios[i].y)
	}

	union := createDisjointUnion(N)

	//Já estão conectadas, então fazemos as uniões iniciais
	var M int
	fmt.Fscan(in, &M)
	for i := 0; i < M; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		union.Union(u, v)
	}

	//Gerar todas as arestas possíveis e calcular as distâncias
	var listaArestas []Aresta
	for i := 1; i <= N; i++ {
		for j := i + 1; j <= N; j++ {
			d := calcularDistancia(predios[i], predios[j])
			listaArestas = append(listaArestas, Aresta{u: i, v: j, dist: d})
		}
	}

	sort.Slice(listaArestas, func(i, j int) bool {
		return listaArestas[i].dist < listaArestas[j].dist
	})

	var custoAdicional float64 = 0
	for _, aresta := range listaArestas {
		if union.Union(aresta.u, aresta.v) {
			custoAdicional += aresta.dist
		}
	}

	fmt.Fprintf(out, "%.3f\n", custoAdicional)
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Ponto struct {
	x, y float64
}

func pontoNoSegmento(p, q, r Ponto) bool {
	if q.x <= max(p.x, r.x) && q.x >= min(p.x, r.x) &&
		q.y <= max(p.y, r.y) && q.y >= min(p.y, r.y) {
		return true
	}
	return false
}

const (
	Colinear = iota
	Horario
	AntiHorario
)

func orientacao(p, q, r Ponto) int {
	val := (q.y-p.y)*(r.x-q.x) - (q.x-p.x)*(r.y-q.y)

	if math.Abs(val) < 1e-9 {
		return int(Colinear) // colinear
	}
	if val > 0 {
		return int(Horario) // horário
	}
	return int(AntiHorario) // anti-horário
}

func intersectam(p1, q1, p2, q2 Ponto) bool {
	p1q1p2 := orientacao(p1, q1, p2)
	p1q1q2 := orientacao(p1, q1, q2)
	p2q2p1 := orientacao(p2, q2, p1)
	p2q2q1 := orientacao(p2, q2, q1)

	// as orientações mudam de sinal para ambos os segmentos
	if p1q1p2 != p1q1q2 && p2q2p1 != p2q2q1 {
		return true
	}

	// quando os pontos são colineares e estão sobre o segmento oposto
	if p1q1p2 == 0 && pontoNoSegmento(p1, p2, q1) {
		return true
	}
	if p1q1q2 == 0 && pontoNoSegmento(p1, q2, q1) {
		return true
	}
	if p2q2p1 == 0 && pontoNoSegmento(p2, p1, q2) {
		return true
	}
	if p2q2q1 == 0 && pontoNoSegmento(p2, q1, q2) {
		return true
	}

	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var p1, q1, p2, q2 Ponto

	if _, err := fmt.Fscan(in, &p1.x, &p1.y, &q1.x, &q1.y); err != nil {
		return
	}
	if _, err := fmt.Fscan(in, &p2.x, &p2.y, &q2.x, &q2.y); err != nil {
		return
	}

	if intersectam(p1, q1, p2, q2) {
		fmt.Fprintln(out, "SIM")
	} else {
		fmt.Fprintln(out, "NAO")
	}
}

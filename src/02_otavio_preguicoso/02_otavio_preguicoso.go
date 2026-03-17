package main

import (
	"fmt"
	"sort"
)

func main() {
	var alturaMin int
	if _, err := fmt.Scan(&alturaMin); err != nil {
		return
	}

	const QNTD_MESES = 12

	alturaMeses := make([]int, QNTD_MESES)
	for i := 0; i < QNTD_MESES; i++ {
		if _, err := fmt.Scan(&alturaMeses[i]); err != nil {
			return
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(alturaMeses)))

	somaAltura := 0
	qntdMesesRega := 0

	for somaAltura < alturaMin && qntdMesesRega < QNTD_MESES {
		somaAltura += alturaMeses[qntdMesesRega]
		qntdMesesRega += 1
	}

	if somaAltura >= alturaMin {
		fmt.Println(qntdMesesRega)
	} else {
		fmt.Println("nao cresce")
	}
}

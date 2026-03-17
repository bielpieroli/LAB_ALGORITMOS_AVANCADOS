package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	if _, err := fmt.Fscan(in, &t); err != nil {
		return
	}

	for tc := 1; tc <= t; tc++ {
		var n int
		if _, err := fmt.Fscan(in, &n); err != nil {
			return
		}

		assuntos := make(map[string]int, n)
		for i := 0; i < n; i++ {
			var linguagem string
			var prazo int
			if _, err := fmt.Fscan(in, &linguagem, &prazo); err != nil {
				return
			}
			assuntos[linguagem] = prazo
		}

		var dias int
		var linguagem string
		if _, err := fmt.Fscan(in, &dias); err != nil {
			return
		}
		if _, err := fmt.Fscan(in, &linguagem); err != nil {
			return
		}

		fmt.Printf("Case %d: ", tc)
		if prazo, ok := assuntos[linguagem]; ok {
			if prazo <= dias {
				fmt.Println("Yessss")
			} else if prazo <= dias+5 {
				fmt.Println("Late")
			} else {
				fmt.Println("Do your own homework!")
			}
		} else {
			fmt.Println("Do your own homework!")
		}
	}
}

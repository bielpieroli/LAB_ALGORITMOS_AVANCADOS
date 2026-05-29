package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var num int
	if _, err := fmt.Fscan(in, &num); err != nil {
		return
	}

	fmt.Fprint(out, num/3, " ", num/5, " ", num/15, "\n")

}

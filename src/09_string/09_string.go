package main

import (
	"bufio"
	"fmt"
	"os"
)

// Indica o maior prefixo que também é sufixo.
func longestPreffixSuffix(pattern string) []int {
	tamPattern := len(pattern)
	longest := make([]int, tamPattern)
	length := 0
	index := 1

	for index < tamPattern {
		if pattern[index] == pattern[length] {
			length++
			longest[index] = length
			index++
		} else {
			if length != 0 {
				length = longest[length-1]
			} else {
				longest[index] = 0
				index++
			}
		}
	}
	return longest
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var str, subStr string
	if _, err := fmt.Fscan(in, &str); err != nil {
		return
	}
	if _, err := fmt.Fscan(in, &subStr); err != nil {
		return
	}

	tamSub := len(subStr)
	tamStr := len(str)

	if tamSub == 0 {
		fmt.Fprintln(out, 0)
		return
	}

	longest := longestPreffixSuffix(subStr)

	nroRepeticoes := 0

	strIndex := 0
	subIndex := 0

	// Percorre a string apenas uma vez O(n)
	for strIndex < tamStr {
		if subStr[subIndex] == str[strIndex] {
			strIndex++
			subIndex++
		}

		// Encontrou uma ocorrência completa
		if subIndex == tamSub {
			nroRepeticoes++

			//Reinicia subIndex usando a tabela para encontrar possíveis sobreposições
			subIndex = longest[subIndex-1]

		} else if strIndex < tamStr && subStr[subIndex] != str[strIndex] {

			// Não encontrou uma completa, por isso, volta do subIndex para o próximo possível início de correspondência
			if subIndex != 0 {
				subIndex = longest[subIndex-1]
			} else {
				strIndex++
			}
		}
	}

	fmt.Fprintln(out, nroRepeticoes)
}

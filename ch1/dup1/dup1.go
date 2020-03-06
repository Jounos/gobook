//DUP1 exibe o testo de toda linha que aparece mais de
//uma vez na entrada-padrao, precedido por sua contagem.package dup1

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	//NOTA: ignorando erros em potencial de input.Err()

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

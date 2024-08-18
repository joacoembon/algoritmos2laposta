package main

import (
	"bufio"
	"fmt"
	"os"
	"tp1/calculadora"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		expresion := scanner.Text()
		resultado := calculadora.Calculadora(expresion)
		fmt.Println(resultado)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stdout, "error:", err)
		os.Exit(1)
	}
}

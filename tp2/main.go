package main

import (
	"bufio"
	"os"
	Acciones "tp2/Acciones"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Acciones.RealizarComandos(scanner)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

func leerArreglo(archivo *os.File) []int {
	scanner := bufio.NewScanner(archivo)
	var arreglo []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error al convertir número:", err)
			os.Exit(1)
		}
		arreglo = append(arreglo, num)
	}
	return arreglo
}

func main() {
	// abre y lee archivo1.in
	archivo1, err := os.Open("archivo1.in")
	if err != nil {
		fmt.Println("Error al abrir archivo1.in:", err)
		return
	}
	defer archivo1.Close()

	arreglo1 := leerArreglo(archivo1)

	// abre y lee archivo2.in
	archivo2, err := os.Open("archivo2.in")
	if err != nil {
		fmt.Println("Error al abrir archivo2.in:", err)
		return
	}
	defer archivo2.Close()

	arreglo2 := leerArreglo(archivo2)

	// compara los arreglos y seleccionar el mayor
	switch ejercicios.Comparar(arreglo1, arreglo2) {
	case -1:
		ejercicios.Seleccion(arreglo2)
		for i := 0; i < len(arreglo2); i++ {
			fmt.Println(arreglo2[i])
		}
	case 0:
		ejercicios.Seleccion(arreglo1)
		for i := 0; i < len(arreglo1); i++ {
			fmt.Println(arreglo1[i])
		}
	case 1:
		ejercicios.Seleccion(arreglo1)
		for i := 0; i < len(arreglo1); i++ {
			fmt.Println(arreglo1[i])
		}
	}
}

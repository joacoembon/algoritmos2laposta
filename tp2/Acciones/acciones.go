package acciones

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	comando "tp2/Comandos"
)

func RealizarComandos(scanner *bufio.Scanner) {
	info := comando.CrearAlmacen()
	for scanner.Scan() {
		linea := scanner.Text()
		linea = strings.TrimSuffix(linea, "\n")
		argumentos := strings.Split(linea, " ")
		comando := argumentos[0]

		switch comando {
		case "agregar_archivo":
			if len(argumentos) != 2 {
				fmt.Fprintf(os.Stderr, "Error en comando %s\n", comando)
				return
			}
			agregar_archivo(info, argumentos[1], comando)

		case "ver_visitantes":
			if len(argumentos) != 3 {
				fmt.Fprintf(os.Stderr, "Error en comando %s\n", comando)
				return
			}
			ver_vistantes(info, argumentos[1], argumentos[2], comando)

		case "ver_mas_visitados":
			if len(argumentos) != 2 {
				fmt.Fprintf(os.Stderr, "Error en comando %s\n", comando)
				return
			}
			ver_mas_visitados(info, argumentos[1], comando)
		default:
			fmt.Fprintf(os.Stderr, "Error en comando %s\n", comando)
			return
		}
	}
	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "Error al leer la entrada estándar:", scanner.Err())
	}
}

func agregar_archivo(info comando.Comandos, nombre_archivo, comando string) {
	err := info.AgregarArchivo(nombre_archivo)
	if !err {
		fmt.Fprintf(os.Stderr, "Error en comando %s\n", comando)
		return
	}
	for _, ip := range info.VerDoS() {
		fmt.Printf("DoS: %s\n", ip)
	}
	fmt.Println("OK")
}

func ver_vistantes(info comando.Comandos, desdeIP, hastaIP, comando string) {
	visitantes := info.VerVisitantes(desdeIP, hastaIP)
	if visitantes == nil {
		fmt.Fprintf(os.Stderr, "Error en comando %s\n", comando)
		return
	}
	fmt.Println("Visitantes:")
	for _, ip := range visitantes {
		fmt.Printf("\t%s\n", ip)
	}
	fmt.Println("OK")
}

func ver_mas_visitados(info comando.Comandos, numString, comando string) {
	num, err := strconv.Atoi(numString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error en comando %s\n", comando)
		return
	}
	mas_visitados := info.VerMasVisitados(num)
	fmt.Fprintf(os.Stdout, "Sitios más visitados:\n")
	for _, sitio := range mas_visitados {
		fmt.Printf("\t%s - %d\n", sitio.URL, sitio.Cant)
	}
	fmt.Println("OK")
}

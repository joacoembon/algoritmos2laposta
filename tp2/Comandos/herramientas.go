package comandos

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	heap "tdas/cola_prioridad"
	Diccionario "tdas/diccionario"
	"time"
)

const (
	SEPARADOR       = "\t"
	VALOR_DEFAULT   = 1
	RANGO_TIEMPO    = 2
	MAX_SOLICITUDES = 5
)

type almacenDatos struct {
	DoS        []string
	Visitantes Diccionario.DiccionarioOrdenado[string, int]
	Sitios     Diccionario.Diccionario[string, int]
}

type entradaLog struct {
	IP     string
	Fecha  time.Time
	Metodo string
	URL    string
}

type urlUsos struct {
	URL  string
	Cant int
}

func CrearAlmacen() Comandos {
	return &almacenDatos{DoS: nil, Visitantes: Diccionario.CrearABB[string, int](maxIP), Sitios: Diccionario.CrearHash[string, int]()}
}

func (info *almacenDatos) AgregarArchivo(nombre_archivo string) bool {
	solicitudesPorIP := Diccionario.CrearHash[string, []time.Time]()
	ipReportadas := Diccionario.CrearHash[string, int]()
	archivo, err := os.Open(nombre_archivo)
	if err != nil {
		return false
	}
	defer archivo.Close()
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		datos := strings.Split(linea, SEPARADOR)
		if len(datos) < 4 {
			continue
		}
		Fecha, err := time.Parse("2006-01-02T15:04:05-07:00", datos[1])
		if err != nil {
			return false
		}
		entrada := entradaLog{IP: datos[0], Fecha: Fecha, Metodo: datos[2], URL: datos[3]}

		if !solicitudesPorIP.Pertenece(entrada.IP) {
			info.Visitantes.Guardar(entrada.IP, VALOR_DEFAULT)
		}
		if info.Sitios.Pertenece(entrada.URL) {
			info.Sitios.Guardar(entrada.URL, info.Sitios.Obtener(entrada.URL)+1)
		} else {
			info.Sitios.Guardar(entrada.URL, 1)
		}
		detectarDoS(entrada, solicitudesPorIP, ipReportadas)
	}
	info.DoS = ordenarIpsRadixSort(ipReportadas)
	return true
}

func (info *almacenDatos) VerDoS() []string {
	return info.DoS
}

func (info *almacenDatos) VerVisitantes(desdeIP, hastaIP string) []string {
	if !validarIp(desdeIP) || !validarIp(hastaIP) {
		return nil
	}
	visitantes := make([]string, 0)
	info.Visitantes.IterarRango(&desdeIP, &hastaIP, func(clave string, dato int) bool {
		visitantes = append(visitantes, clave)
		return true
	})
	return visitantes
}

func (info *almacenDatos) VerMasVisitados(k int) []urlUsos {
	arr := make([]urlUsos, 0)
	cantVisitados := 0
	info.Sitios.Iterar(func(clave string, dato int) bool {
		Url := urlUsos{URL: clave, Cant: dato}
		arr = append(arr, Url)
		cantVisitados++
		return true
	})
	kPrimeros := make([]urlUsos, minimo(k, cantVisitados))
	heap := heap.CrearHeapArr(arr, funcionCmpMaximo)

	for i := range minimo(k, cantVisitados) {
		kPrimeros[i] = heap.Desencolar()
	}
	return kPrimeros
}

func detectarDoS(entrada entradaLog, solicitudesPorIP Diccionario.Diccionario[string, []time.Time], ipReportadas Diccionario.Diccionario[string, int]) {
	ip := entrada.IP
	if ipReportadas.Pertenece(ip) {
		return
	}
	if solicitudesPorIP.Pertenece(ip) {
		solicitudesPorIP.Guardar(ip, append(solicitudesPorIP.Obtener(ip), entrada.Fecha))
	} else {
		solicitudesPorIP.Guardar(ip, []time.Time{entrada.Fecha})
	}
	cantidad_solicitudes := len(solicitudesPorIP.Obtener(ip))
	if cantidad_solicitudes >= MAX_SOLICITUDES {
		if entrada.Fecha.Sub(solicitudesPorIP.Obtener(ip)[cantidad_solicitudes-MAX_SOLICITUDES]).Seconds() < RANGO_TIEMPO {
			ipReportadas.Guardar(ip, VALOR_DEFAULT)
		}
	}
}

func ordenarIpsRadixSort(ipReportadas Diccionario.Diccionario[string, int]) []string {
	arr := make([]string, ipReportadas.Cantidad())
	cont := 0
	ipReportadas.Iterar(func(ip string, dato int) bool {
		arr[cont] = ip
		cont++
		return true
	})
	countingSort(arr, 256, 3, convertir)
	countingSort(arr, 256, 2, convertir)
	countingSort(arr, 256, 1, convertir)
	countingSort(arr, 256, 0, convertir)
	return arr
}

func countingSort(elementos []string, rango, digitos int, convertir func(string) int) {
	frecuencias := make([]int, rango)
	sumasAcumuladas := make([]int, rango)
	resultados := make([]string, len(elementos))

	for _, elem := range elementos {
		valor := convertir(strings.Split(elem, ".")[digitos])
		frecuencias[valor]++
	}
	for i := 1; i < len(frecuencias); i++ {
		sumasAcumuladas[i] = sumasAcumuladas[i-1] + frecuencias[i-1]
	}
	for _, elem := range elementos {
		valor := convertir(strings.Split(elem, ".")[digitos])
		pos := sumasAcumuladas[valor]
		resultados[pos] = elem
		sumasAcumuladas[valor]++
	}
	copy(elementos, resultados)
}

func convertir(num string) int {
	numInt, _ := strconv.Atoi(num)
	return numInt
}

func validarIp(ip string) bool {
	numerosIp := strings.Split(ip, ".")
	if len(numerosIp) != 4 {
		return false
	}
	for _, numero := range numerosIp {
		numeroInt, err := strconv.Atoi(numero)
		if err != nil || numeroInt < 0 || numeroInt > 255 {
			return false
		}
	}
	return true
}

func funcionCmpMaximo(a, b urlUsos) int {
	return a.Cant - b.Cant
}

func minimo(num1, num2 int) int {
	if num1 <= num2 {
		return num1
	}
	return num2
}

func maxIP(ip1, ip2 string) int {
	segmento1 := strings.Split(ip1, ".")
	segmento2 := strings.Split(ip2, ".")

	for i := 0; i < 4; i++ {
		num1, _ := strconv.Atoi(segmento1[i])
		num2, _ := strconv.Atoi(segmento2[i])
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}
	return 0
}

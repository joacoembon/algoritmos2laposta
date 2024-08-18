package ejercicios

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	i := *x
	*x = *y
	*y = i
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	if len(vector) == 0 {
		return -1
	}
	posMax := 0

	for i := 0; i < len(vector); i++ {
		if vector[i] > vector[posMax] {
			posMax = i
		}
	}
	return posMax
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	for i := 0; i < len(vector1) && i < len(vector2); i++ {

		if vector1[i] < vector2[i] {
			return -1
		}

		if vector1[i] > vector2[i] {
			return 1
		}
	}

	if len(vector1) < len(vector2) {
		return -1
	} else if len(vector1) > len(vector2) {
		return 1
	}

	return 0
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	for i := len(vector) - 1; i >= 0; i-- {
		// encuentra la posición del máximo elemento en el subarreglo no ordenado
		max := Maximo(vector[:i+1])
		// intercambia el elemento
		Swap(&vector[i], &vector[max])
	}
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	return sumaW(vector, 0)
}

func sumaW(vector []int, indice int) int {
	if len(vector) == 0 || indice == len(vector) {
		return 0
	}
	return vector[indice] + sumaW(vector, indice+1)
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func EsCadenaCapicua(cadena string) bool {
	return esCadenaCapicuaW(cadena, 0, len(cadena)-1)
}

func esCadenaCapicuaW(cadena string, inicio int, fin int) bool {
	if inicio >= fin {
		return true
	}
	if cadena[inicio] != cadena[fin] {
		return false
	}
	return esCadenaCapicuaW(cadena, inicio+1, fin-1)
}

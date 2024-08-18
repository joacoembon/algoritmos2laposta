package cola_prioridad_test

import (
	"cmp"
	"math/rand"
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

var FUNCION_INT = cmp.Compare[int]

func TestColaVacia(t *testing.T) {
	t.Log("Comprueba que una cola vacia se comporte como tal")
	heap := TDAHeap.CrearHeap(FUNCION_INT)
	require.True(t, heap.Cantidad() == 0)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() }, "No se puede ver el maximo de una cola")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() }, "No se puede desencolar de una cola vacia")
}

func TestColaConElementosNoEstaVacia(t *testing.T) {
	t.Log("Comprueba el correcto funcionamiento de la primitiva EstaVacia()")
	heap := TDAHeap.CrearHeap(FUNCION_INT)
	heap.Encolar(5)
	require.False(t, heap.EstaVacia())
}

func TestFuncionamientoCorrecto(t *testing.T) {
	t.Log("Comprueba el correcto funcionamiento de la cola para uno y más elementos")
	heap := TDAHeap.CrearHeap(FUNCION_INT)
	heap.Encolar(1)
	require.Equal(t, 1, heap.VerMax())
	require.True(t, heap.Cantidad() == 1)
	heap.Encolar(0)
	require.Equal(t, 1, heap.VerMax())
	heap.Encolar(7)
	heap.Encolar(5)
	heap.Encolar(3)
	require.Equal(t, 7, heap.VerMax())
	require.True(t, heap.Cantidad() == 5)
}

func TestDesencolar(t *testing.T) {
	t.Log("Comprueba el correcto funcionamiento de la primitiva Desencolar() y que la cola se " +
		"comporte correctamente en todo momento")
	heap := TDAHeap.CrearHeap(FUNCION_INT)
	heap.Encolar(5)
	heap.Encolar(4)
	heap.Encolar(3)
	heap.Encolar(2)
	heap.Encolar(1)
	heap.Encolar(0)
	require.Equal(t, 5, heap.VerMax())
	require.Equal(t, 6, heap.Cantidad())
	for i := 5; !heap.EstaVacia(); i-- {
		require.Equal(t, i, heap.Desencolar())
		require.True(t, heap.Cantidad() == i)
	}
}

func TestVolumen(t *testing.T) {
	t.Log("Comprueba el correcto funcionamiento de la cola si manejamos una gran cantidad de elementos")
	heap := TDAHeap.CrearHeap(FUNCION_INT)
	arreglo := make([]int, 10000)
	for i := range arreglo {
		arreglo[i] = rand.Intn(100)
	}
	for i := range arreglo {
		require.Equal(t, i, heap.Cantidad())
		heap.Encolar(arreglo[i])
	}
	max := heap.VerMax()
	for i := 9999; !heap.EstaVacia(); i-- {
		require.GreaterOrEqual(t, max, heap.Desencolar())
		require.True(t, heap.Cantidad() == i)
	}
	require.Equal(t, 0, heap.Cantidad())
}

func TestCrearHeapDesdeArreglo(t *testing.T) {
	t.Log("Comprueba el correcto funcionamiento de la primitiva CrearHeapArr()" +
		"y que luego se comporte como un heap")
	heap := TDAHeap.CrearHeapArr([]int{5, 0, 3, 2, 7, 1, 6}, FUNCION_INT)
	require.True(t, heap.Cantidad() == 7)
	require.Equal(t, 7, heap.VerMax())
	heap.Encolar(8)
	heap.Encolar(7)
	max := heap.VerMax()
	for i := 8; !heap.EstaVacia(); i-- {
		require.GreaterOrEqual(t, max, heap.Desencolar())
		require.True(t, heap.Cantidad() == i)
	}
}

func TestCrearHeapDesdeArregloVacio(t *testing.T) {
	t.Log("Comprueba el correcto funcionamiento de la primitiva CrearHeapArr() con arreglo vacio" +
		"y que luego se comporte como un heap")
	heap := TDAHeap.CrearHeapArr([]int{}, FUNCION_INT)
	require.True(t, heap.Cantidad() == 0)
	heap.Encolar(8)
	heap.Encolar(6)
	require.True(t, heap.Cantidad() == 2)
	require.Equal(t, 8, heap.VerMax())
}

func TestVolumenCrearHeapDesdeArreglo(t *testing.T) {
	t.Log("Comprueba el correcto funcionamiento del ordenamiento HeapSort con un arreglo grande")
	arreglo := make([]int, 10000)
	for i := range arreglo {
		arreglo[i] = rand.Intn(100)
	}
	heap := TDAHeap.CrearHeapArr(arreglo, FUNCION_INT)
	heap.Encolar(101)
	require.Equal(t, heap.VerMax(), 101)
	require.Equal(t, heap.Cantidad(), 10001)
	heap.Desencolar()
	heap.Desencolar()
	heap.Desencolar()
	require.Equal(t, heap.Cantidad(), 9998)
	max := heap.VerMax()
	for i := 9997; !heap.EstaVacia(); i-- {
		require.GreaterOrEqual(t, max, heap.Desencolar())
		require.True(t, heap.Cantidad() == i)
	}
}

func TestHeapSort(t *testing.T) {
	t.Log("Comprueba el correcto funcionamiento del ordenamiento HeapSort")
	arreglo := []int{7, 4, 9, 5, 8, 3, 2, 0, 6, 4}
	TDAHeap.HeapSort(arreglo, FUNCION_INT)
	require.Equal(t, arreglo, []int{0, 2, 3, 4, 4, 5, 6, 7, 8, 9})
}

func TestVolumenHeapSort(t *testing.T) {
	t.Log("Comprueba el correcto funcionamiento del ordenamiento HeapSort con un arreglo grande")
	arreglo := make([]int, 10000)
	for i := range arreglo {
		arreglo[i] = rand.Intn(100)
	}
	TDAHeap.HeapSort(arreglo, FUNCION_INT)
	for i := 0; i < len(arreglo)-1; i++ {
		require.LessOrEqual(t, arreglo[i], arreglo[i+1])
	}
}

func TestHeapSortArregloVacio(t *testing.T) {
	t.Log("Comprueba el correcto funcionamiento del ordenamiento HeapSort en un arreglo vacio")
	arreglo := []int{}
	TDAHeap.HeapSort(arreglo, FUNCION_INT)
	require.Equal(t, arreglo, []int{})
}

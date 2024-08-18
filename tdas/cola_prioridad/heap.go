package cola_prioridad

const (
	CAPACIDADINICIAL = 5
)

type heap[T any] struct {
	tabla    []T
	cantidad int
	cmp      func(T, T) int
}

func CrearHeap[T any](funcionCmp func(T, T) int) ColaPrioridad[T] {
	tabla := make([]T, CAPACIDADINICIAL)
	return &heap[T]{tabla: tabla, cantidad: 0, cmp: funcionCmp}
}

func CrearHeapArr[T any](arreglo []T, funcionCmp func(T, T) int) ColaPrioridad[T] {
	arregloNuevo := make([]T, len(arreglo))
	copy(arregloNuevo, arreglo)
	heap := &heap[T]{tabla: arregloNuevo, cantidad: len(arreglo), cmp: funcionCmp}
	heapify(arregloNuevo, len(arregloNuevo), funcionCmp)
	if cap(heap.tabla) == 0 {
		redimension(&heap.tabla, CAPACIDADINICIAL)
	}
	return heap
}

func HeapSort[T any](arreglo []T, funcionCmp func(T, T) int) {
	heapify(arreglo, len(arreglo), funcionCmp)
	ordenar(arreglo, len(arreglo), funcionCmp)
}

func ordenar[T any](arreglo []T, cantidad int, funcionCmp func(T, T) int) {
	for i := len(arreglo) - 1; i >= 0; i-- {
		swap(&arreglo[i], &arreglo[0])
		cantidad--
		downheap(0, arreglo, cantidad, funcionCmp)
	}
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) Encolar(elemento T) {
	if heap.cantidad == cap(heap.tabla) {
		redimension(&heap.tabla, cap(heap.tabla)*2)
	}
	heap.tabla[heap.cantidad] = elemento
	upheap(heap.cantidad, heap.tabla, heap.cmp)
	heap.cantidad++
}

func (heap *heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.tabla[0]
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	if heap.cantidad < cap(heap.tabla)/4 {
		redimension(&heap.tabla, cap(heap.tabla)/2)
	}
	desencolado := heap.tabla[0]
	swap(&heap.tabla[0], &heap.tabla[heap.cantidad-1])
	heap.cantidad--
	downheap(0, heap.tabla, heap.cantidad, heap.cmp)
	return desencolado
}

func (heap *heap[T]) Cantidad() int {
	return heap.cantidad
}

func swap[T any](x *T, y *T) {
	*x, *y = *y, *x
}

func upheap[T any](posicion int, arreglo []T, funcCmp func(T, T) int) {
	if (posicion-1)/2 < 0 {
		return
	}
	posicionPadre := (posicion - 1) / 2
	comparacion := funcCmp(arreglo[posicion], arreglo[posicionPadre])
	if comparacion > 0 {
		swap(&arreglo[posicion], &arreglo[posicionPadre])
		upheap(posicionPadre, arreglo, funcCmp)
	}
}

func downheap[T any](posicion int, arreglo []T, cantidad int, funcCmp func(T, T) int) {
	posicionHijoMayor := calcularHijoMayor(posicion, arreglo, cantidad, funcCmp)
	if posicionHijoMayor < 0 {
		return
	}
	comparacion := funcCmp(arreglo[posicion], arreglo[posicionHijoMayor])
	if comparacion >= 0 {
		return
	}
	swap(&arreglo[posicion], &arreglo[posicionHijoMayor])
	downheap(posicionHijoMayor, arreglo, cantidad, funcCmp)

}

func calcularHijoMayor[T any](posicion int, arreglo []T, cantidad int, funcCmp func(T, T) int) int {
	if posicion >= (cantidad / 2) {
		return -1
	}
	if posicion < (cantidad/2)-1 || cantidad%2 != 0 {
		comparacion := funcCmp(arreglo[(2*posicion)+1], arreglo[(2*posicion)+2])
		if comparacion < 0 {
			return (2 * posicion) + 2
		}
	}
	return (2 * posicion) + 1
}

func heapify[T any](arreglo []T, cantidad int, funcionCmp func(T, T) int) {
	for i := cantidad - 1; i >= 0; i-- {
		downheap(i, arreglo, cantidad, funcionCmp)
	}
}

func redimension[T any](arreglo *[]T, tamArreglo int) {
	colaNueva := make([]T, tamArreglo)
	copy(colaNueva, *arreglo)
	*arreglo = colaNueva
}

package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}

}

func crearNodoLista[T any](dato T, proximo *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato: dato, siguiente: proximo}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil

}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevoNodo := crearNodoLista(dato, lista.primero)
	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	}
	lista.primero = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevoNodo := crearNodoLista(dato, nil)
	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}
	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	if lista.largo == 1 {
		lista.ultimo = nil
	}
	primero := lista.primero
	lista.primero = primero.siguiente
	lista.largo--
	return primero.dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato

}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato

}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo

}

func (lista listaEnlazada[T]) Iterar(visitar func(T) bool) {
	if lista.EstaVacia() {
		return
	}
	nodoActual := lista.primero
	for nodoActual != nil && visitar(nodoActual.dato) {
		nodoActual = nodoActual.siguiente
	}
}

//Iterador Externo:

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{actual: lista.primero, anterior: nil, lista: lista}

}

func (iter *iterListaEnlazada[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")

	}
	return iter.actual.dato
}

func (iter *iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil

}

func (iter *iterListaEnlazada[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")

	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iterListaEnlazada[T]) Insertar(dato T) {
	nuevoNodo := crearNodoLista(dato, nil)

	if iter.anterior == nil {
		// Caso borde: Insertar al principio
		nuevoNodo.siguiente = iter.actual
		iter.actual = nuevoNodo
		iter.lista.primero = nuevoNodo
		if iter.lista.largo == 0 {
			iter.lista.ultimo = nuevoNodo
		}
	} else if iter.actual == nil {
		// Caso borde: Insetar a lo ultimo
		iter.anterior.siguiente = nuevoNodo
		nuevoNodo.siguiente = iter.actual
		iter.actual = nuevoNodo
		iter.lista.ultimo = nuevoNodo
	} else {
		// Caso normal
		iter.anterior.siguiente = nuevoNodo
		nuevoNodo.siguiente = iter.actual
		iter.actual = nuevoNodo
	}
	iter.lista.largo++

}
func (iter *iterListaEnlazada[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")

	}
	valor := iter.actual.dato

	if iter.anterior == nil {
		// Caso borde: borrar primero
		siguiente := iter.actual.siguiente
		iter.actual.siguiente = nil
		iter.actual = siguiente
		iter.lista.primero = iter.actual

	} else if iter.actual.siguiente == nil {
		// Caso borde: borrar ultimo
		iter.anterior.siguiente = nil
		iter.actual = nil
		iter.lista.ultimo = iter.anterior

	} else {
		iter.anterior.siguiente = iter.actual.siguiente
		iter.actual.siguiente = nil
		iter.actual = iter.anterior.siguiente

	}
	iter.lista.largo--

	return valor
}

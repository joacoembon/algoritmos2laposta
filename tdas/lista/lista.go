package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos, falso en caso contrario.
	EstaVacia() bool

	// InsertarPrimero inserta un elemento en la primera posicion de la lista
	InsertarPrimero(T)

	// InsertarUltimo inserta un elemento en la primera posicion de la lista
	InsertarUltimo(T)

	// BorrarPrimero elimina el primer elemento de la lista.
	// Si esta vacia, entra en panico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero devuelve el primer elemento de la lista.
	// Si esta vacia, entra en panico con un mensaje "La lista esta vacia".
	VerPrimero() T

	// VerUltimo devuelve el ultimo elemento de la lista.
	// Si esta vacia, entra en panico con un mensaje "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos de la lista
	Largo() int

	// Iterar recorre la lista y ejecuta la funcion visitar para cada elemento de la lista.
	// Si la funcion visitar devuelve falso, o la lista no tiene más elementos, se corta la iteracion.
	// El usuario debe asegurarse que la funcion de visita coincida lo que recibe por parametro con el tipo de elementos en la lista. En caso contrario, no compila.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un iterador
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual devuelve el elemento de la posicion en la que el iterador esta parado de la lista
	// Si el iterador ya termino de iterar, entra en panico con un mensaje "El iterador termino de iterar"
	VerActual() T

	// HaySiguiente devuelve verdadero si hay un elemento siguiente para ver, falso en caso contrario.
	HaySiguiente() bool

	// Siguiente avanza a la siguiente posicion en la lista
	// Si el iterador ya termino de iterar, entra en panico con un mensaje "El iterador termino de iterar"
	Siguiente()

	// Insertar inserta un elemento en la posicion en la que el iterador esta parado de la lista.
	Insertar(T)

	// Borrar elimina el elemento de la posicion en la que el iterador esta parado de la lista.
	// Si el iterador ya termino de iterar, entra en panico con un mensaje "El iterador termino de iterar"
	Borrar() T
}
